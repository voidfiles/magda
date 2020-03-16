package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/voidfiles/magda/graph"
	"github.com/voidfiles/magda/pkg/repository"
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
	logger := server.MustNewLogger()

	authClient, err := app.Auth(context)
	if err != nil {
		logger.Fatal().AnErr("error", err).Msg("Failed to initialize auth client")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	firestoreClient, err := app.Firestore(context)
	if err != nil {
		logger.Fatal().AnErr("error", err).Msg("Failed to get firestore client")
	}
	repo := repository.MustNewRepository(firestoreClient)
	resolver := graph.MustNewResolver(repo)
	finalSrv := server.BuildGraphQLServer(logger, authClient, resolver)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", finalSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
