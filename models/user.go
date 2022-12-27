package models

type User struct {
	ID    int    `json:"id"`
	FIO   string `json:"fio"`
	Age   int    `json:"age"`
	Login string `json:"login"`
	Pin   string `json:"pin"`
}
