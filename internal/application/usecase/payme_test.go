package usecase_test

import (
	"errors"
	"github.com/m3k3r1/payme/internal/application/usecase"
	"github.com/m3k3r1/payme/internal/domain"
	"github.com/m3k3r1/payme/internal/infra/memory"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PaymeTestSuite struct {
	testAccount     *domain.Account
	paymeusecase    *usecase.PaymeUseCase
	paymeRepository *memory.InMemoryPaymeKeyRepository
	suite.Suite
}

func (suite *PaymeTestSuite) SetupSuite() {
	suite.testAccount, _ = domain.NewAccount(
		"test",
		"test",
		"test",
		"000-00",
	)

	suite.paymeRepository = memory.NewInMemoryPaymeKeyRepository()
	suite.paymeRepository.SaveAccount(suite.testAccount)

	suite.paymeusecase = &usecase.PaymeUseCase{
		PaymeKeyRepository: suite.paymeRepository,
	}
}

func (suite *PaymeTestSuite) TestRegisterKey() {
	paymeKey, err := suite.paymeusecase.RegisterKey("000-00", suite.testAccount.ID)
	suite.Nil(err)
	suite.NotNil(paymeKey)

	key, err := suite.paymeRepository.FindKey(paymeKey.Key)
	suite.Nil(err)
	suite.Equal(paymeKey.Key, key.Key)
}

func (suite *PaymeTestSuite) TestRegisterKey_Error() {
	type testCase struct {
		name        string
		key         string
		accountID   string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "account not found",
			key:         "000-00",
			accountID:   "asdf",
			expectedErr: domain.ErrAccountNotFound,
		},
		{
			name:        "key is empty",
			key:         "",
			accountID:   suite.testAccount.ID,
			expectedErr: domain.ErrPaymeKeyEmpty,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			_, err := suite.paymeusecase.RegisterKey(tc.key, tc.accountID)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func (suite *PaymeTestSuite) TestFindKey() {
	paymeKey, err := suite.paymeusecase.RegisterKey("000-00", suite.testAccount.ID)
	suite.Nil(err)
	suite.NotNil(paymeKey)

	key, err := suite.paymeusecase.FindKey(paymeKey.Key)
	suite.Nil(err)
	suite.Equal(paymeKey.Key, key.Key)
}

func (suite *PaymeTestSuite) TestFindKey_Error() {
	key, err := suite.paymeusecase.FindKey("asdfasd")
	suite.Nil(key)
	suite.Equal(domain.ErrPaymeKeyNotFound, err)
}

func TestPaymeTestSuite(t *testing.T) {
	suite.Run(t, new(PaymeTestSuite))
}
