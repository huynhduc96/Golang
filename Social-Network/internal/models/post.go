package models

type Post struct {
	Id               int    `json:"id,omitempty"`
	ContentText      string `json:"text,omitempty"`
	ContentImagePath string `json:"imagePath"`
	CreatedAt        string `json:"createdAt,omitempty"`
	UserId           int    `json:"userId,omitempty"`
	DownloadUrl      string `json:"downloadUrl,omitempty"`
	Visible          int    `json:"visible,omitempty"`
}

type CreatePost struct {
	ContentText string
}
