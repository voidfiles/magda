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
	"github.com/rs/zerolog/hlog"
	"github.com/voidfiles/magda/graph"
	"github.com/voidfiles/magda/graph/generated"
	"github.com/voidfiles/magda/pkg/server"
	"google.golang.org/api/option"
)

const defaultPort = "8080"

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
	logger, err := server.MustNewLogger()
	if err != nil {
		log.Fatalf("error initializing logger: %v", err)
	}

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
	loggerMiddleware := hlog.RequestHandler("request")
	logSrv := loggerMiddleware(http.HandlerFunc(authSrv.ServeHTTP))
	finalSrv := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := zerolog.Ctx(r.Context())
		log.Info().Msg("Authorized Request")
		logSrv.ServeHTTP(w, r)
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", finalSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
