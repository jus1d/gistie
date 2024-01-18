package gist

type CreateRequest struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type CreateResponse struct {
	URL string `json:"html_url"`
}

type File struct {
	Content string `json:"content"`
}
