package domain

import (
	"errors"
	"time"
)

var (
	ErrIDEmpty           = errors.New("id is empty")
	ErrBankIDEmpty       = errors.New("bank id is empty")
	ErrAccountOwnerEmpty = errors.New("account owner is empty")
	ErrAccountPhoneEmpty = errors.New("account phone is empty")
	ErrAccountNotFound   = errors.New("account not found")
)

type Account struct {
	ID          string    `json:"id"`
	BankID      string    `json:"bank_id"`
	Owner       string    `json:"owner"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewAccount(accountID string, bankID string, owner string, phoneNumber string) (*Account, error) {
	account := Account{
		ID:          accountID,
		BankID:      bankID,
		Owner:       owner,
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (account *Account) Validate() error {
	if account.ID == "" {
		return ErrIDEmpty
	}
	if account.BankID == "" {
		return ErrBankIDEmpty
	}
	if account.Owner == "" {
		return ErrAccountOwnerEmpty
	}
	if account.PhoneNumber == "" {
		return ErrAccountPhoneEmpty
	}
	return nil
}
