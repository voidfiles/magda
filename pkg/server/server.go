package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rs/zerolog"
	"github.com/voidfiles/magda/graph"
	"github.com/voidfiles/magda/graph/generated"
)

// RequestHandler adds the request method and URL as a field to the context's logger
// using fieldKey as field key.
func RequestHandler() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := zerolog.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("method", r.Method).Str("path", r.URL.String())
			})
			next.ServeHTTP(w, r)
		})
	}
}

// BuildGraphQLServer constructs A GraphQL Server
func BuildGraphQLServer(logger zerolog.Logger, authClient authClient, resolver *graph.Resolver) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	authSrv := MustNewAuthMiddleware(srv, &logger, authClient)
	logSrv := RequestHandler()(http.HandlerFunc(authSrv.ServeHTTP))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localLogger := logger.With().Str("yo", "whats up").Logger()
		r = r.WithContext(
			localLogger.WithContext(
				r.Context(),
			),
		)
		logSrv.ServeHTTP(w, r)
	})
}
