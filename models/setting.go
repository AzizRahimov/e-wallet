package models

import "time"

type Settings struct {
	AppParams      Params           `json:"app"`
	PostgresParams PostgresSettings `json:"postgres_params"`
	SecretKey      SecretKey        `json:"secretKey"`
	Auth           Auth             `json:"auth"`
}

type Params struct {
	ServerName string `json:"server_name"`
	PortRun    string `json:"port_run"`
	LogFile    string `json:"log_file"`
	ServerURL  string `json:"server_url"`
}

type PostgresSettings struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	DataBase string `json:"database"`
}

type SecretKey struct {
	Key string `json:"key"`
}

const (
	salt       = "dkaskdsa21312das3das"
	signingKey = "quiche#SKDJASDKA3dsa213#sH"
	tokenTTL   = 12 * time.Hour
)

type Auth struct {
	Salt       string `json:"salt"`
	SigningKey string `json:"signing_key"`
}
