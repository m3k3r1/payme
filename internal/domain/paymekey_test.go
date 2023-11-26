package domain_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/m3k3r1/payme/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymekey_NewPaymeKey(t *testing.T) {
	key := "test"
	account, _ := domain.NewAccount(uuid.New().String(), "test", "test", "test")
	paymeKey, err := domain.NewPaymeKey(key, account)
	assert.Nil(t, err)
	assert.NotNil(t, paymeKey)
	assert.Equal(t, key, paymeKey.Key)
	assert.Equal(t, account, paymeKey.Account)
	assert.Equal(t, "active", paymeKey.Status)
}

func TestPaymekey_NewPaymeKey_Error(t *testing.T) {
	type testCase struct {
		name        string
		key         string
		account     *domain.Account
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "key is empty",
			key:         "",
			account:     &domain.Account{},
			expectedErr: domain.ErrPaymeKeyEmpty,
		},
		{
			name:        "account is empty",
			key:         "test",
			account:     nil,
			expectedErr: domain.ErrPaymeKeyAccountEmpty,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			paymeKey, err := domain.NewPaymeKey(tc.key, tc.account)
			assert.Nil(t, paymeKey)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
