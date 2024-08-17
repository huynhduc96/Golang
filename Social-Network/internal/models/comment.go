package models

type Comment struct {
	Id          int    `json:"id"`
	PostId      int    `json:"postId"`
	UserId      int    `json:"userId"`
	ContentText string `json:"content"`
	CreatedAt   string `json:"createdAt,omitempty"`
}

type CreateComment struct {
	ContentText string
}
