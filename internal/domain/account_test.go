package domain_test

import (
	"github.com/google/uuid"
	"github.com/m3k3r1/payme/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount_NewAccount(t *testing.T) {
	accountID := uuid.New().String()
	bankID := "test"
	owner := "test"
	phoneNumber := "test"
	account, err := domain.NewAccount(accountID, bankID, owner, phoneNumber)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, bankID, account.BankID)
	assert.Equal(t, owner, account.Owner)
	assert.Equal(t, phoneNumber, account.PhoneNumber)
}

func TestAccount_NewAccount_Error(t *testing.T) {
	type testCase struct {
		name        string
		accountID   string
		bankID      string
		owner       string
		phoneNumber string
		expectedErr error
	}

	testCases := []testCase{
		{
			accountID:   "",
			name:        "account id is empty",
			bankID:      "test",
			owner:       "test",
			phoneNumber: "test",
			expectedErr: domain.ErrIDEmpty,
		},
		{
			accountID:   "test",
			name:        "bank id is empty",
			bankID:      "",
			owner:       "test",
			phoneNumber: "test",
			expectedErr: domain.ErrBankIDEmpty,
		},
		{
			accountID:   "test",
			name:        "owner is empty",
			bankID:      uuid.New().String(),
			owner:       "",
			phoneNumber: "test",
			expectedErr: domain.ErrAccountOwnerEmpty,
		},
		{
			accountID:   "test",
			name:        "phone is empty",
			bankID:      uuid.New().String(),
			owner:       "test",
			phoneNumber: "",
			expectedErr: domain.ErrAccountPhoneEmpty,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			account, err := domain.NewAccount(tc.accountID, tc.bankID, tc.owner, tc.phoneNumber)
			assert.Nil(t, account)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
