package shortener

import (
	"errors"
	"net/url"
)

var errURLFileSchema = errors.New("schema 'file://' not allowed")
var errURLInvalid = errors.New("invalid URL")

var slashByte = byte(47)
var maxLengthURL = 2048

// NormalizeURL validates and normalizes given rawURL string.
func NormalizeURL(rawURL []byte) ([]byte, error) {
	url, err := normalizeURLPassOne(rawURL)
	if err != nil {
		return []byte{}, err
	}

	url, err = normalizeURLPassTwo(url)
	if err != nil {
		return []byte{}, err
	}

	return url, nil
}

func normalizeURLPassOne(rawURL []byte) ([]byte, error) {
	if len(rawURL) > maxLengthURL {
		return []byte{}, errURLInvalid
	}

	u, err := url.Parse(string(rawURL))
	if err != nil {
		return []byte{}, errURLInvalid
	}

	if u.Scheme == "" {
		u.Scheme = "http"
	}

	if u.Scheme == "file" {
		return []byte{}, errURLFileSchema
	}

	if u.Host == "" && (u.Path == "" || u.Path[0] == slashByte) {
		return []byte{}, errURLInvalid
	}

	return []byte(u.String()), nil
}

func normalizeURLPassTwo(rawURL []byte) ([]byte, error) {
	u, err := url.Parse(string(rawURL))
	if err != nil {
		return []byte{}, errURLInvalid
	}

	if u.Host != "" && u.Path == "" {
		u.Path = "/"
	}

	return []byte(u.String()), nil
}
