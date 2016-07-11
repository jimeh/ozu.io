package web

//go:generate easyjson -all responses.go

// ShortenedResponse contains shortened URL info.
type ShortenedResponse struct {
	UID    string `json:"uid"`
	URL    string `json:"url"`
	Target string `json:"target"`
}

// ErrorResponse contains error info.
type ErrorResponse struct {
	Error string `json:"error"`
}
