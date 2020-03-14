package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type WebsiteKind string

const (
	// site is the default website value
	WebsiteKindSite      WebsiteKind = "site"
	WebsiteKindWikimedia WebsiteKind = "wikimedia"
	WebsiteKindArtnet    WebsiteKind = "artnet"
	WebsiteKindArtsy     WebsiteKind = "artsy"
	// homepage is for the main online presence of an entity
	WebsiteKindHomepage WebsiteKind = "homepage"
)

var AllWebsiteKind = []WebsiteKind{
	WebsiteKindSite,
	WebsiteKindWikimedia,
	WebsiteKindArtnet,
	WebsiteKindArtsy,
	WebsiteKindHomepage,
}

func (e WebsiteKind) IsValid() bool {
	switch e {
	case WebsiteKindSite, WebsiteKindWikimedia, WebsiteKindArtnet, WebsiteKindArtsy, WebsiteKindHomepage:
		return true
	}
	return false
}

func (e WebsiteKind) String() string {
	return string(e)
}

func (e *WebsiteKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WebsiteKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WebsiteKind", str)
	}
	return nil
}

func (e WebsiteKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// A website represents a page on the world wide web. It's rich text and a jumping off place for more information about something.
//
// Websites are notoriously temporary. The goal should be to hook into some kind of persistent identifier scheme.
//
// But, they are a nescesity. Things like homepages, and blogs of artists however temporary are important, but should be kept up to date.
type Website struct {
	ID          string      `json:"id"`
	URL         string      `json:"url"`
	Kind        WebsiteKind `json:"kind"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
