package http_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/rs/zerolog"
	"github.com/steinfletcher/apitest"
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

func buildGraphQLQuery(operationName, query string, variables interface{}) string {
	data := map[string]interface{}{
		"operationName": operationName,
		"variables":     variables,
		"query":         query,
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return string(b)
}

func TestGraphQL(t *testing.T) {
	qr := buildGraphQLQuery(
		"createWebsite",
		`query createWebsite($input: WebsiteInput!) {
			createWebsite(input: $input) {
				id
				url
			}
		}`,
		map[string]interface{}{
			"input": map[string]string{
				"url":         "http://google.com",
				"kind":        "site",
				"title":       "Google",
				"description": "A search interface",
			},
		},
	)

	h, token := handler()
	apitest.New().
		Handler(h).
		Debug().
		Post("/").
		Header("Authorization", fmt.Sprintf("Bearer %s", token)).
		Header("Content-Type", "application/json").
		Body(qr).
		Expect(t).
		Body(`{"data":{"createWebsite":{"id":"aa2239c17609b21eba034c564af878f3eec8ce83ed0f2768597d2bc2fd4e4da5","url":"http://google.com"}}}`).
		Status(http.StatusOK).
		End()
}
