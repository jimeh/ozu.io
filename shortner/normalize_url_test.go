package shortner

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var examples = []struct {
	valid      bool
	url        string
	normalized string
	error      string
}{
	{valid: true, url: "google.com", normalized: "http://google.com/"},
	{valid: true, url: "google.com/", normalized: "http://google.com/"},
	{valid: true, url: "http://google.com", normalized: "http://google.com/"},
	{valid: true, url: "http://google.com/"},
	{valid: true, url: "https://google.com", normalized: "https://google.com/"},
	{valid: true, url: "https://google.com/"},
	{valid: true, url: "google.yeah", normalized: "http://google.yeah/"},
	{valid: true, url: "http://news.google.com/"},

	{valid: true, url: "http://google.com/?h=en&foo=bar"},
	{valid: true,
		url:        "http://google.com?h=en&foo=bar",
		normalized: "http://google.com/?h=en&foo=bar"},
	{valid: true,
		url:        "google.com/?h=en&foo=bar",
		normalized: "http://google.com/?h=en&foo=bar"},
	{valid: true,
		url:        "google.com?h=en&foo=bar",
		normalized: "http://google.com/?h=en&foo=bar"},

	{valid: true, url: "http://google.com/#nope"},
	{valid: true,
		url:        "http://google.com#nope",
		normalized: "http://google.com/#nope"},
	{valid: true,
		url:        "google.com/#nope",
		normalized: "http://google.com/#nope"},
	{valid: true,
		url:        "google.com#nope",
		normalized: "http://google.com/#nope"},

	{valid: true, url: "http://google.com/?h=en&foo=bar#nope"},
	{valid: true,
		url:        "http://google.com?h=en&foo=bar#nope",
		normalized: "http://google.com/?h=en&foo=bar#nope"},
	{valid: true,
		url:        "google.com/?h=en&foo=bar#nope",
		normalized: "http://google.com/?h=en&foo=bar#nope"},
	{valid: true,
		url:        "google.com?h=en&foo=bar#nope",
		normalized: "http://google.com/?h=en&foo=bar#nope"},

	{valid: true, url: "(248034)", normalized: "http://(248034)/"},
	{
		valid: false,
		url:   "*$)]+_<?)",
		error: "invalid URL",
	},
	{
		valid: false,
		url:   "",
		error: "invalid URL",
	},
	{
		valid: false,
		url:   "file:///bin/bash",
		error: "schema 'file://' not allowed",
	},
	{
		valid: false,
		url:   "/users/view.php?uid=138495",
		error: "invalid URL",
	},
	{
		valid:      true,
		url:        "users/view.php?uid=138495",
		normalized: "http://users/view.php?uid=138495",
	},
	{
		valid: true,
		url:   "http://long.com/" + strings.Repeat("0", 2032),
	},
	{
		valid: false,
		url:   "http://long.com/" + strings.Repeat("0", 3000),
		error: "invalid URL",
	},
}

func TestNormalizeURL(t *testing.T) {
	assert := assert.New(t)

	for _, e := range examples {
		result, err := NormalizeURL([]byte(e.url))

		if e.valid {
			assert.Nil(err)
			if e.normalized != "" {
				assert.Equal([]byte(e.normalized), result)
			} else {
				assert.Equal([]byte(e.url), result)
			}
		} else {
			assert.NotNil(err, "Expected error, got nil.")
			if e.error != "" {
				assert.EqualError(err, e.error, "URL: "+e.url)
			}
		}
	}
}
