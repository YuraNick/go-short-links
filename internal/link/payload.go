package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinkReadRequest struct {
	Hash string `json:"hash" validate:"required"`
}

type LinkReadResponse struct {
	Url string `json:"url"`
}
