package graph

import (
	"github.com/voidfiles/magda/pkg/repository"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver connectes schema resolvers to graphql server
type Resolver struct {
	repo repository.Repository
}

// MustNewResolver creates a resolver
func MustNewResolver(repo repository.Repository) *Resolver {
	return &Resolver{
		repo: repo,
	}
}
