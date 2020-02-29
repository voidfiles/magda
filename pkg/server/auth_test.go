package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type TestAuthClient struct {
	token *auth.Token
	err   error
}

func (tac TestAuthClient) VerifyIDToken(ctx context.Context, token string) (*auth.Token, error) {
	return tac.token, tac.err
}

func TestAuthMiddleware(t *testing.T) {
	nl := zerolog.Nop()
	// nl := zerolog.New(os.Stdout)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var uid string
		var ok bool
		if uid, ok = r.Context().Value(UIDContextKey).(string); !ok {
			http.Error(w, "nope", http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`"%s"`, uid)))
		}

	})

	tests := []struct {
		authToken            *auth.Token
		authError            error
		expectedStatusCode   int
		authHeader           string
		expectedResponseBody string
	}{
		{
			authToken:            &auth.Token{UID: "abc"},
			authError:            nil,
			expectedStatusCode:   http.StatusOK,
			authHeader:           "Bearer abc",
			expectedResponseBody: `"abc"`,
		},
		{
			authToken:            &auth.Token{UID: "abc"},
			authError:            nil,
			expectedStatusCode:   http.StatusBadRequest,
			authHeader:           "bbb abc",
			expectedResponseBody: `{"error": {"code": "missing-token", "message": "Authorization token must follow bearer scheme"}}`,
		},
		{
			authToken:            &auth.Token{UID: "abc"},
			authError:            nil,
			expectedStatusCode:   http.StatusBadRequest,
			authHeader:           "bbb abc zzz",
			expectedResponseBody: `{"error": {"code": "missing-token", "message": "Authorization token must follow bearer scheme"}}`,
		},
		{
			authToken:            &auth.Token{UID: "abc"},
			authError:            fmt.Errorf("Failed to verify token"),
			expectedStatusCode:   http.StatusUnauthorized,
			authHeader:           "Bearer abc",
			expectedResponseBody: `{"error": {"code": "unauthorized", "message": "Token is invalid"}}`,
		},
	}

	for _, tt := range tests {
		tac := TestAuthClient{
			token: tt.authToken,
			err:   tt.authError,
		}
		stack := MustNewAuthMiddleware(handler, &nl, tac)
		r := httptest.NewRequest("GET", "/", strings.NewReader(""))
		r.Header.Add("Authorization", tt.authHeader)
		w := httptest.NewRecorder()

		stack.ServeHTTP(w, r)

		assert.Equal(t, tt.expectedStatusCode, w.Code)
		assert.JSONEq(t, tt.expectedResponseBody, w.Body.String())
	}

}
