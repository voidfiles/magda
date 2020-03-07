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

func MustNewURLizer(schemes []string) URLizer {
	return URLizer{
		schemes: validSchemes(schemes),
	}
}

type URLizer struct {
	schemes validSchemes
}

func (u URLizer) Validate(rawurl string) (string, error) {
	rawurl = strings.Trim(rawurl, " ")
	parsedUrl, err := urlx.ParseWithDefaultScheme(rawurl, "https")

	if err != nil {
		return rawurl, errors.Wrapf(err, "Failed to parse url: %s", rawurl)
	}

	if !u.schemes.IsValid(parsedUrl.Scheme) {
		return rawurl, fmt.Errorf("scheme: %s is invalid must be one of %s", parsedUrl.Scheme, u.schemes.String())
	}

	if parsedUrl.RawPath != "" {
		parsedUrl.RawPath = strings.Trim(parsedUrl.RawPath, "/")
	}

	validatedURL, err := urlx.Normalize(parsedUrl)

	if err != nil {
		return rawurl, fmt.Errorf("Failed to normalize %s", rawurl)
	}

	return validatedURL, nil
}
