package web

//go:generate easyjson -all response.go

// Response contains shortened URL info.
type Response struct {
	UID    string `json:"uid"`
	URL    string `json:"url"`
	Target string `json:"target"`
	Error  string `json:"error"`
}
