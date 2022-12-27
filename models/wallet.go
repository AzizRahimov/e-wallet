package models

type Wallet struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	Balance      float64 `json:"balance"`
	Account      string  `json:"account"`
	IsIdentified bool    `json:"is_identified"`
}
