package models

type Like struct {
	PostId int `json:"postId,omitempty"`
	UserId int `json:"userId,omitempty"`
}

type LikeActionPost struct {
	PostId int
}
