package dto

import "time"

type App struct {
	Id                 int
	Name               string
	EncryptedSecretKey string
	PublicKey          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
