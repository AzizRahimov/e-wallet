package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Status    string    `json:"status"`
	TransType string    `json:"trans_type"`
	CreatedAt time.Time `json:"created_at"`
}
