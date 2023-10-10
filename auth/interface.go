package auth

import (
	"time"

	"github.com/o1egl/paseto"
)

type Maker interface {
	CreateToken(data map[string]string) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}
