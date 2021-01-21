package main

type WebsocketEvent struct {
	Type  string  `json:"type"`
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Debt  int     `json:"debt"`
	Price float64 `json:"price"`
	User  User    `json:"user"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Debt int    `json:"debt"`
}

type OpenBarData struct {
	Price float64 `json:"price"`
	Users []User  `json:"users"`
}
