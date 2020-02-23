package main

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

type ContextKey string

const (
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
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		am.logger.Error().Msgf("Request missing authorization header")
		http.Error(w, `{"error": {"code": "missing-header", "message": "Missing authorization header"}}`, http.StatusUnauthorized)
		return
	}

	authParts := strings.Fields(authHeader)

	if strings.ToLower(authParts[0]) != "bearer" {
		am.logger.Error().Msgf("Authorization header missing bearer schema: %v", authHeader)
		http.Error(w, `{"error": {"code": "missing-token", "message": "Authorization token must follow bearer scheme"}}`, http.StatusBadRequest)
		return
	}

	if len(authParts) > 2 {
		am.logger.Error().Msgf("Authorization header has too many parts %v", authHeader)
		http.Error(w, `{"error": {"code": "missing-token", "message": "Authorization token must follow bearer scheme"}}`, http.StatusBadRequest)
		return
	}
	ctx := context.TODO()
	token, err := am.client.VerifyIDToken(ctx, authParts[1])
	if err != nil {
		am.logger.Fatal().AnErr("error", err).Msg("error verifying ID token:")
	}

	childLogger := am.logger.With().Str("uid", token.UID).Logger()
	ctx = childLogger.WithContext(ctx)
	ctx = context.WithValue(ctx, UIDContextKey, token.UID)
	r = r.WithContext(ctx)

	am.handler.ServeHTTP(w, r)
}
