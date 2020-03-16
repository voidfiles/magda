package repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/voidfiles/magda/graph/model"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

type token struct {
	token string
}

func (t token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + t.token,
	}, nil
}

func (t token) RequireTransportSecurity() bool {
	return false
}

// CreateInsecureJWT creates an token for use in testing
func CreateInsecureJWT(uid, role string) token {
	header := `{"alg":"none","kid":"fakekid","typ":"JWT"}`
	body := fmt.Sprintf(`{"iat":%d,"eat":%d,"sub":"%s","uid":"%s","role":"%s"}`, time.Now().Unix(), time.Now().Unix()+10000, uid, uid, role)

	t := fmt.Sprintf(
		"%s.%s.%s",
		base64.RawURLEncoding.EncodeToString([]byte(header)),
		base64.RawURLEncoding.EncodeToString([]byte(body)),
		"",
	)

	return token{token: t}
}

func getClient() *firestore.Client {
	ctx := context.Background()

	token := CreateInsecureJWT("ab", "admin")
	conn, err := grpc.Dial(
		"localhost:8972",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(token),
	)
	if err != nil {
		log.Fatal(err)
	}
	client, err := firestore.NewClient(
		ctx,
		"magdatest",
		option.WithGRPCConn(conn),
	)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
func TestCreateWebsiteError(t *testing.T) {
	client := getClient()
	r := MustNewRepository(client)

	hello := "hello"
	description := "this is a great website"

	cases := []struct {
		input  model.WebsiteInput
		output string
	}{
		{
			input: model.WebsiteInput{
				URL:         "ftp://google.com",
				Kind:        model.WebsiteKindSite,
				Title:       &hello,
				Description: &description,
			},
			output: "scheme: ftp is invalid must be one of http, https",
		},
	}

	for _, c := range cases {
		_, err := r.CreateWebsite(context.TODO(), c.input)
		assert.EqualError(t, err, c.output)
	}

}

func refOfString(str string) *string {
	return &str
}

func TestCreateWebsite(t *testing.T) {
	client := getClient()
	r := MustNewRepository(client)

	cases := []struct {
		input  model.WebsiteInput
		output model.Website
	}{
		{
			input: model.WebsiteInput{
				URL:         "https://example.com",
				Kind:        model.WebsiteKindSite,
				Title:       refOfString("Hello"),
				Description: refOfString("Description"),
			},
			output: model.Website{
				URL:         "https://example.com",
				Kind:        model.WebsiteKindSite,
				Title:       refOfString("Hello"),
				Description: refOfString("Description"),
			},
		},
	}

	for _, ce := range cases {
		website, err := r.CreateWebsite(context.TODO(), ce.input)
		assert.NoError(t, err)
		assert.Equal(t, ce.output.URL, website.URL)
		assert.Equal(t, ce.input.Kind, website.Kind)
		assert.True(t, website.Kind.IsValid())
		assert.Equal(t, ce.input.Title, website.Title)
		assert.Equal(t, ce.input.Description, website.Description)
		assert.NotNil(t, website.CreatedAt)
		assert.NotNil(t, website.UpdatedAt)
		assert.NotNil(t, website.ID)
	}
}

func TestFindWebsite(t *testing.T) {
	client := getClient()
	r := MustNewRepository(client)

	cases := []struct {
		input  model.WebsiteSearch
		output model.Website
	}{
		{
			input: model.WebsiteSearch{
				ID:  refOfString("100680ad546ce6a577f42f52df33b4cfdca756859e664b8d7de329b150d09ce9"),
				URL: nil,
			},
			output: model.Website{
				ID:  "100680ad546ce6a577f42f52df33b4cfdca756859e664b8d7de329b150d09ce9",
				URL: "https://example.com",
			},
		},
		{
			input: model.WebsiteSearch{
				ID:  refOfString("100680ad546ce6a577f42f52df33b4cfdca756859e664b8d7de329b150d09ce9"),
				URL: refOfString("https://example.com"),
			},
			output: model.Website{
				ID:  "100680ad546ce6a577f42f52df33b4cfdca756859e664b8d7de329b150d09ce9",
				URL: "https://example.com",
			},
		},
		{
			input: model.WebsiteSearch{
				ID:  nil,
				URL: refOfString("https://example.com"),
			},
			output: model.Website{
				ID:  "100680ad546ce6a577f42f52df33b4cfdca756859e664b8d7de329b150d09ce9",
				URL: "https://example.com",
			},
		},
	}
	_, err := r.CreateWebsite(context.TODO(), model.WebsiteInput{
		URL:  "https://example.com",
		Kind: model.WebsiteKindSite,
	})
	assert.NoError(t, err)
	for _, ce := range cases {
		website, err := r.FindWebsite(context.TODO(), ce.input)
		assert.NoError(t, err)
		assert.Equal(t, website.ID, ce.output.ID)

	}
}
