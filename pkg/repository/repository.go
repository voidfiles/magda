package repository

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

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

type createWebsite struct {
	ID          string    `firestore:"id"`
	URL         string    `firestore:"url"`
	Title       *string   `firestore:"title"`
	Description *string   `firestore:"description"`
	Kind        string    `firestore:"kind"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

func (r repository) createWebsite(ctx context.Context, website model.WebsiteInput) (model.Website, error) {
	url, err := r.urlizer.Validate(website.URL)

	if err != nil {
		return model.Website{}, err
	}

	websites := r.fs.Collection("Websites")
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	websiteRef := websites.Doc(id)

	forCreate := createWebsite{
		ID:          id,
		URL:         url,
		Title:       website.Title,
		Description: website.Description,
		Kind:        website.Kind.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = websiteRef.Create(ctx, forCreate)
	if err != nil {
		if status.Code(err) != codes.AlreadyExists {
			return model.Website{}, errors.Wrap(err, "Failed save website")
		}
	}

	wr, err := websiteRef.Get(ctx)
	if err != nil {
		return model.Website{}, errors.Wrap(err, "Failed save website")
	}

	var cw createWebsite

	err = wr.DataTo(&cw)

	if err != nil {
		return model.Website{}, errors.Wrap(err, "Failed save website ab")
	}

	newWebsite := model.Website{
		URL:         cw.URL,
		Title:       cw.Title,
		Description: cw.Description,
		Kind:        website.Kind,
		CreatedAt:   cw.CreatedAt,
		UpdatedAt:   cw.UpdatedAt,
	}

	return newWebsite, nil
}
