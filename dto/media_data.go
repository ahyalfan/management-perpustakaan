package dto

type MediaData struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

type CreatedMediaRequest struct {
	Path string `json:"path"`
}
