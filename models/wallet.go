package models

type Wallet struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	Balance      float64 `json:"balance"`
	Account      string  `json:"account"`
	IsIdentified bool    `json:"is_identified"`
}

type TopUp struct {
	ClientID      int     `json:"client_id"`
	ReceiverPhone string  `json:"receiver_phone"`
	Balance       float64 `json:"balance"`
}
