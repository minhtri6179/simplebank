package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	Expired  time.Time `json:"expired"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:       tokenID,
		Username: username,
		IssuedAt: time.Now(),
		Expired:  time.Now().Add(duration),
	}
	return payload, nil
}
func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expired) {
		return ErrExpiredToken
	}
	return nil
}
