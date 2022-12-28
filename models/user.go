package models

type User struct {
	ID    int    `json:"id,omitempty"`
	FIO   string `json:"fio,omitempty"`
	Age   int    `json:"age,omitempty"`
	Phone string `json:"phone,omitempty"`
	Pin   string `json:"pin,omitempty"`
}
