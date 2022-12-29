package models

type Wallet struct {
	ID           int     `json:"id,"`
	UserID       int     `json:"user_id"`
	Balance      float64 `json:"balance"`
	Account      string  `json:"account,omitempty"`
	IsIdentified bool    `json:"is_identified"`
}

type TopUp struct {
	ClientID      int    `json:"client_id,omitempty"`
	ReceiverPhone string `json:"receiver_phone"`

	Amount float64 `json:"amount,omitempty"`
}
