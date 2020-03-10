package repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/voidfiles/magda/graph/model"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

func createInsecureJWT(uid, role string) string {
	header := `{"alg":"none","kid":"fakekid"}`
	body := fmt.Sprintf(`{"iat":0,"sub":"%s","uid":"%s","role":"%s"}`, uid, uid, role)

	return fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString([]byte(header)),
		base64.RawURLEncoding.EncodeToString([]byte(body)),
	)
}

func getClient() *firestore.Client {
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8972")
	ctx := context.Background()

	token := createInsecureJWT("a", "admin")
	fmt.Printf("%s\n", token)
	client, err := firestore.NewClient(
		ctx,
		"magdatest",
		option.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken:  createInsecureJWT("a", "admin"),
			TokenType:    "Bearer",
			RefreshToken: "",
		})),
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
		_, err := r.createWebsite(context.TODO(), c.input)
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
		website, err := r.createWebsite(context.TODO(), ce.input)
		assert.NoError(t, err)
		assert.Equal(t, ce.output.URL, website.URL)
		// assert.Equal(t, websiteInput.Kind, website.Kind)
		// assert.True(t, website.Kind.IsValid())
		// assert.Equal(t, websiteInput.Title, website.Title)
		// assert.Equal(t, websiteInput.Description, website.Description)
		// assert.NotNil(t, website.CreatedAt)
		// assert.NotNil(t, website.UpdatedAt)
		// assert.NotNil(t, website.ID)
	}
}
