package repository

import (
	"fmt"
	"strings"

	"github.com/goware/urlx"
	"github.com/pkg/errors"
)

type validSchemes []string

func (v validSchemes) IsValid(scheme string) bool {
	for _, validScheme := range v {
		if scheme == validScheme {
			return true
		}
	}

	return false
}

func (v validSchemes) String() string {
	return strings.Join(v, ", ")
}

// MustNewURLizer Creates a new URLizer
func MustNewURLizer(schemes []string) URLizer {
	return URLizer{
		schemes: validSchemes(schemes),
	}
}

// URLizer validates and normalizes a URL
type URLizer struct {
	schemes validSchemes
}

// Validate validates and normalizes a url
func (u URLizer) Validate(rawurl string) (string, error) {
	rawurl = strings.Trim(rawurl, " ")
	parsedURL, err := urlx.ParseWithDefaultScheme(rawurl, "https")

	if err != nil {
		return rawurl, errors.Wrapf(err, "Failed to parse url: %s", rawurl)
	}

	if !u.schemes.IsValid(parsedURL.Scheme) {
		return rawurl, fmt.Errorf("scheme: %s is invalid must be one of %s", parsedURL.Scheme, u.schemes.String())
	}

	if parsedURL.RawPath != "" {
		parsedURL.RawPath = strings.Trim(parsedURL.RawPath, "/")
	}

	validatedURL, err := urlx.Normalize(parsedURL)

	if err != nil {
		return rawurl, fmt.Errorf("Failed to normalize %s", rawurl)
	}

	return validatedURL, nil
}
