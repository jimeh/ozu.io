package web

//go:generate easyjson -all responses.go

// URLResponse contains shortened URL info.
type URLResponse struct {
	UID    string `json:"uid"`
	URL    string `json:"url"`
	Target string `json:"target"`
}

// ErrorJSONResponse contains error info.
type ErrorResponse struct {
	Error string `json:"error"`
}
