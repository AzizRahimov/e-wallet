package models

import "time"

type Transaction struct {
	ID          int       `json:"id,omitempty"`
	FromPhone   string    `json:"from_phone,omitempty"`
	ToPhone     string    `json:"to_phone,omitempty"`
	Status      string    `json:"status,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
	TotalAmount float64   `json:"total_amount,omitempty"`
	TrnType     string    `json:"trn_type,omitempty"`
	Operation   int       `json:"operation,omitempty"`
	Month       string    `json:"month,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
