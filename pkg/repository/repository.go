package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/voidfiles/magda/graph/model"
)

// Repository is the interface to magda data
type Repository interface {
	createWebsite(website model.WebsiteInput) (model.Website, error)
}

type firestoreClient interface {
	Collection(string) *firestore.CollectionRef
}

// MustNewRepository creates a new repository
func MustNewRepository() Repository {
	return repository{
		urlizer: MustNewURLizer([]string{"http", "https"}),
	}
}

type repository struct {
	urlizer URLizer
}

func (r repository) createWebsite(website model.WebsiteInput) (model.Website, error) {
	url, err := r.urlizer.Validate(website.URL)

	if err != nil {
		return model.Website{}, err
	}

	return model.Website{
		URL: url,
	}, nil
}
