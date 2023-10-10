package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)



var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
	symmetricKey = os.Getenv("TOKEN_SECRET")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}



func (maker *PasetoMaker) CreateToken(data map[string]string) (string, error) {
	payload := Payload{
		Id:        data["id"],
		Email:     data["email"],
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(1 * time.Hour),
	}

	if len(symmetricKey) != chacha20poly1305.KeySize {
		return "", fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker = &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
