package models

type User struct {
	ID       int    `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Exp      int64  `json:"exp"`
}

type JWT struct {
	Token string `json:"token"`
}
