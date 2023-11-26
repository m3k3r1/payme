package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrPaymeKeyEmpty        = errors.New("payme key is empty")
	ErrPaymeKeyAccountEmpty = errors.New("payme key account is empty")
	ErrPaymeKeyNotFound     = errors.New("payme key not found")
)

type PaymeKeyRepository interface {
	SaveAccount(account *Account) error
	FindAccountByID(id string) (*Account, error)
	SavePaymeKey(paymeKey *PaymeKey) error
	FindKey(key string) (*PaymeKey, error)
}

type PaymeKey struct {
	ID        string    `json:"id"`
	Kind      string    `json:"kind"`
	Key       string    `json:"key"`
	Account   *Account  `json:"account"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPaymeKey(key string, account *Account) (*PaymeKey, error) {
	p := PaymeKey{
		ID:        uuid.New().String(),
		Kind:      "phone",
		Key:       key,
		Account:   account,
		Status:    "active",
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *PaymeKey) Validate() error {
	if p.Key == "" {
		return ErrPaymeKeyEmpty
	}
	if p.Account == nil {
		return ErrPaymeKeyAccountEmpty
	}
	return nil
}
