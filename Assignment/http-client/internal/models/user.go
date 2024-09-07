package models

// User struct to define the user object
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address" `
	UserName string `json:"username" `
	Password string `json:"password" `
}

type UserLogin struct {
	UserName string
	Password string
}
