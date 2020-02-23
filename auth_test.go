package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
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
	//nl := zerolog.Nop()
	nl := zerolog.New(os.Stdout)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var uid string
		var ok bool
		if uid, ok = r.Context().Value(UIDContextKey).(string); !ok {
			http.Error(w, "nope", http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(uid))
		}

	})

	tac := TestAuthClient{
		token: &auth.Token{UID: "abc"},
		err:   nil,
	}
	stack := MustNewAuthMiddleware(handler, &nl, tac)
	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
	r.Header.Add("Authorization", "Bearer abc")
	w := httptest.NewRecorder()

	stack.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "abc", w.Body.String())
}
