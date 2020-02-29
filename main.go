package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog"
	"github.com/voidfiles/magda/graph"
	"github.com/voidfiles/magda/graph/generated"
	"github.com/voidfiles/magda/pkg/server"
	"google.golang.org/api/option"
)

const defaultPort = "8080"

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

func main() {

	cfg, err := server.ReadConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	opt := option.WithCredentialsFile(cfg.GoogleApplicationCredentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	context := context.Background()
	logger := server.MustNewLogger()

	authClient, err := app.Auth(context)
	if err != nil {
		logger.Fatal().AnErr("error", err).Msg("Failed to initialize auth client")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	authSrv := server.MustNewAuthMiddleware(srv, &logger, authClient)
	logSrv := RequestHandler()(http.HandlerFunc(authSrv.ServeHTTP))
	finalSrv := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localLogger := logger.With().Str("yo", "whats up").Logger()
		r = r.WithContext(
			localLogger.WithContext(
				r.Context(),
			),
		)
		logSrv.ServeHTTP(w, r)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", finalSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
