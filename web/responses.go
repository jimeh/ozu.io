package web

//go:generate easyjson -all responses.go

// ShortenedResponse contains shortened URL info.
type ShortenedResponse struct {
	UID      string `json:"uid"`
	ShortURL string `json:"short_url"`
	URL      string `json:"url"`
}

// ErrorResponse contains error info.
type ErrorResponse struct {
	Error string `json:"error"`
}
