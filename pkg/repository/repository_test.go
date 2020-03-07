package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voidfiles/magda/graph/model"
)

func TestCreateWebsiteError(t *testing.T) {
	r := MustNewRepository()

	hello := "hello"
	description := "this is a great website"

	cases := []struct {
		input  model.WebsiteInput
		output string
	}{
		{
			input: model.WebsiteInput{
				URL:         "ftp://google.com",
				Kind:        model.WebsiteKindSite,
				Title:       &hello,
				Description: &description,
			},
			output: "scheme: ftp is invalid must be one of http, https",
		},
	}

	for _, c := range cases {
		_, err := r.createWebsite(c.input)
		assert.EqualError(t, err, c.output)
	}

}

func refOfString(str string) *string {
	return &str
}

func TestCreateWebsite(t *testing.T) {
	r := MustNewRepository()

	cases := []struct {
		input  model.WebsiteInput
		output model.Website
	}{
		{
			input: model.WebsiteInput{
				URL:         "https://example.com",
				Kind:        model.WebsiteKindSite,
				Title:       refOfString("Hello"),
				Description: refOfString("Description"),
			},
			output: model.Website{
				URL:         "https://example.com",
				Kind:        model.WebsiteKindSite,
				Title:       refOfString("Hello"),
				Description: refOfString("Description"),
			},
		},
	}

	for _, ce := range cases {
		website, err := r.createWebsite(ce.input)
		assert.NoError(t, err)
		assert.Equal(t, ce.output.URL, website.URL)
		// assert.Equal(t, websiteInput.Kind, website.Kind)
		// assert.True(t, website.Kind.IsValid())
		// assert.Equal(t, websiteInput.Title, website.Title)
		// assert.Equal(t, websiteInput.Description, website.Description)
		// assert.NotNil(t, website.CreatedAt)
		// assert.NotNil(t, website.UpdatedAt)
		// assert.NotNil(t, website.ID)
	}
}
