package models

import (
	"encoding/gob"
)

type User struct {
	Email        string
	AccessToken  string
	RefreshToken string
	ProviderName string
}

func init() {
	gob.Register(&User{})
}
