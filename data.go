package main

// TODO payload optimization:
// each field gets send with
// the its default value type
// single every time
type WebsocketEvent struct {
	Type  string      `json:"type"`
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Debt  int         `json:"debt"`
	Price float64     `json:"price"`
	User  User        `json:"user"`
	All   OpenBarData `json:"all"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Debt int    `json:"debt"`
}

type OpenBarData struct {
	Price float64 `json:"price"`
	Users []User  `json:"users"`
}
