package http_test

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/brianvoe/gofakeit/v4"
	"github.com/rs/zerolog"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/voidfiles/magda/graph"
	"github.com/voidfiles/magda/pkg/repository"
	"github.com/voidfiles/magda/pkg/server"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

type authClient struct {
	tokenValid bool
	UID        string
}

func (ac authClient) VerifyIDToken(context.Context, string) (*auth.Token, error) {
	if ac.tokenValid {
		return &auth.Token{}, errors.New("fail")
	}

	return &auth.Token{
		UID: ac.UID,
	}, nil
}

type token struct {
	token string
}

func (t token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + t.token,
	}, nil
}

func (t token) Token() string {
	return t.token
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

func handler() (http.HandlerFunc, string) {
	gofakeit.Seed(time.Now().UnixNano())

	logger := zerolog.New(os.Stdout)
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
	repo := repository.MustNewRepository(client)
	resolver := graph.MustNewResolver(repo)
	ac := &authClient{
		tokenValid: false,
		UID:        "ab",
	}

	return server.BuildGraphQLServer(logger, ac, resolver), token.Token()
}

type Interaction struct {
	operationName string
	query         string
	variables     map[string]interface{}
	response      string
	status        int
	assertions    []apitest.Assert
}

func (i Interaction) buildGraphQLQuery() string {
	data := map[string]interface{}{
		"operationName": i.operationName,
		"variables":     i.variables,
		"query":         i.query,
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return string(b)
}

func validateInteraction(t *testing.T, i Interaction) {
	h, token := handler()
	request := apitest.New().
		Handler(h).
		Debug().
		Post("/").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		Header("Content-Type", "application/json").
		Body(i.buildGraphQLQuery())

	response := request.
		Expect(t)

	for _, as := range i.assertions {
		response = response.Assert(as)
	}

	response.
		Status(http.StatusOK).
		End()
}

func matchISO(path string) apitest.Assert {
	return jsonpath.Matches(path, `^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(\\.[0-9]+)?(Z)?$`)
}

func hashURL(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
}

func TestGraphQL(t *testing.T) {

	url := gofakeit.URL()
	title := gofakeit.Word()
	description := gofakeit.Sentence(5)

	is := []Interaction{
		{
			operationName: "createWebsite",
			query: `query createWebsite($input: WebsiteInput!) {
				createWebsite(input: $input) {
					id
					url
					title
					description
					createdAt
					updatedAt
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]string{
					"url":         url,
					"kind":        "site",
					"title":       title,
					"description": description,
				},
			},
			assertions: []apitest.Assert{
				jsonpath.Matches("$.data.createWebsite.id", `^\w+$`),
				jsonpath.Equal("$.data.createWebsite.url", url),
				jsonpath.Equal("$.data.createWebsite.title", title),
				jsonpath.Equal("$.data.createWebsite.description", description),
				matchISO("$.data.createWebsite.createdAt"),
				matchISO("$.data.createWebsite.updatedAt"),
			},
			status: http.StatusOK,
		},
		{
			operationName: "findWebsite",
			query: `query findWebsite($input: WebsiteSearch!) {
				findWebsite(input: $input) {
					id
					url
					title
					description
					createdAt
					updatedAt
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]string{
					"id": hashURL(url),
				},
			},
			assertions: []apitest.Assert{
				jsonpath.Matches("$.data.findWebsite.id", `^\w+$`),
				jsonpath.Equal("$.data.findWebsite.url", url),
				jsonpath.Equal("$.data.findWebsite.title", title),
				jsonpath.Equal("$.data.findWebsite.description", description),
				matchISO("$.data.findWebsite.createdAt"),
				matchISO("$.data.findWebsite.updatedAt"),
			},
			status: http.StatusOK,
		},
		{
			operationName: "findWebsite",
			query: `query findWebsite($input: WebsiteSearch!) {
				findWebsite(input: $input) {
					id
					url
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]string{
					"url": url,
				},
			},
			assertions: []apitest.Assert{
				jsonpath.Matches("$.data.findWebsite.id", `^\w+$`),
				jsonpath.Equal("$.data.findWebsite.url", url),
			},
			status: http.StatusOK,
		},
		{
			operationName: "findWebsite",
			query: `query findWebsite($input: WebsiteSearch!) {
				findWebsite(input: $input) {
					id
					url
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]string{
					"url": strings.ToUpper(url),
				},
			},
			assertions: []apitest.Assert{
				jsonpath.Matches("$.data.findWebsite.id", `^\w+$`),
				jsonpath.Equal("$.data.findWebsite.url", url),
			},
			status: http.StatusOK,
		},
	}
	for _, i := range is {
		validateInteraction(t, i)
	}

}
