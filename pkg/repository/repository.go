package repository

import (
	"context"
	"crypto/sha256"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/voidfiles/magda/graph/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Repository is the interface to magda data
type Repository interface {
	createWebsite(ctx context.Context, website model.WebsiteInput) (model.Website, error)
}

// MustNewRepository creates a new repository
func MustNewRepository(fs *firestore.Client) Repository {
	return repository{
		urlizer: MustNewURLizer([]string{"http", "https"}),
		fs:      fs,
	}
}

type repository struct {
	urlizer URLizer
	fs      *firestore.Client
}

func (r repository) createWebsite(ctx context.Context, website model.WebsiteInput) (model.Website, error) {
	url, err := r.urlizer.Validate(website.URL)

	if err != nil {
		return model.Website{}, err
	}

	// TODO: I need to add in server time stamp thingies here
	// TODO: Also need to make sure to maybe unhook the create model, from model.Website
	newWebsite := model.Website{
		URL:         url,
		Title:       website.Title,
		Description: website.Description,
	}

	websites := r.fs.Collection("Websites")
	websiteRef := websites.Doc(fmt.Sprintf("%x", sha256.Sum256([]byte(url))))

	_, err = websiteRef.Create(ctx, newWebsite)
	if err != nil {
		if status.Code(err) != codes.AlreadyExists {
			return model.Website{}, errors.Wrap(err, "Failed save website")
		}
	}

	return newWebsite, nil
}
