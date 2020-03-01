package server

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/rs/zerolog"
)

type authClient interface {
	VerifyIDToken(context.Context, string) (*auth.Token, error)
}

// ContextKey is for string collision detection
type ContextKey string

const (
	//UIDContextKey is the specific ContextKey used for the user unique id
	UIDContextKey ContextKey = "UID"
)

// AuthMiddleware authenticates http requests
type AuthMiddleware struct {
	client  authClient
	logger  *zerolog.Logger
	handler http.Handler
}

// MustNewAuthMiddleware creates an auth http middleware
func MustNewAuthMiddleware(handler http.Handler, logger *zerolog.Logger, client authClient) AuthMiddleware {
	return AuthMiddleware{
		client:  client,
		logger:  logger,
		handler: handler,
	}
}

func (am AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		logger.Error().Msgf("Request missing authorization header")
		http.Error(w, `{"error": {"code": "missing-header", "message": "Missing authorization header"}}`, http.StatusUnauthorized)
		return
	}

	authParts := strings.Fields(authHeader)

	if strings.ToLower(authParts[0]) != "bearer" {
		logger.Error().Msgf("Authorization header missing bearer schema: %v", authHeader)
		http.Error(w, `{"error": {"code": "missing-token", "message": "Authorization token must follow bearer scheme"}}`, http.StatusBadRequest)
		return
	}
	if len(authParts) != 2 {
		logger.Error().Msgf("Authorization header has too many parts %v", authHeader)
		http.Error(w, `{"error": {"code": "missing-token", "message": "Authorization token must follow bearer scheme"}}`, http.StatusBadRequest)
		return
	}

	token, err := am.client.VerifyIDToken(r.Context(), authParts[1])
	if err != nil {
		logger.Error().AnErr("error", err).Msg("error verifying ID token:")
		http.Error(w, `{"error": {"code": "unauthorized", "message": "Token is invalid"}}`, http.StatusUnauthorized)
		return
	}

	childLogger := logger.With().Str("uid", token.UID).Logger()

	r = r.WithContext(
		context.WithValue(
			childLogger.WithContext(r.Context()),
			UIDContextKey,
			token.UID,
		),
	)
	childLogger.Log().Msg("request start")
	am.handler.ServeHTTP(w, r)
}
