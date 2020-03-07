package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "ftp://google.com",
			output: "scheme: ftp is invalid must be one of https",
		},
		{
			input:  "ftp://",
			output: "Failed to parse url: ftp://: host : empty host",
		},
		{
			input:  "****",
			output: "Failed to parse url: ****: host ****: invalid host",
		},
	}

	urlizer := MustNewURLizer([]string{"https"})
	for _, ce := range cases {
		_, err := urlizer.Validate(ce.input)
		assert.EqualError(t, err, ce.output)
	}
}

func TestValidation(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "HTTPS://google.com ",
			output: "https://google.com",
		},
		{
			input:  "https://google.com?b=c&a=d",
			output: "https://google.com?a=d&b=c",
		},
	}

	urlizer := MustNewURLizer([]string{"https"})
	for _, ce := range cases {
		out, err := urlizer.Validate(ce.input)
		assert.NoError(t, err)
		assert.Equal(t, ce.output, out)
	}
}
