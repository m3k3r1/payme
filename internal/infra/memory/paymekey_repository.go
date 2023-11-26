package memory

import (
	"github.com/m3k3r1/payme/internal/domain"
	"sync"
)

type InMemoryPaymeKeyRepository struct {
	paymeKeys map[string]*domain.PaymeKey
	accounts  map[string]*domain.Account
	sync.Mutex
}

func NewInMemoryPaymeKeyRepository() *InMemoryPaymeKeyRepository {
	return &InMemoryPaymeKeyRepository{
		paymeKeys: make(map[string]*domain.PaymeKey),
		accounts:  make(map[string]*domain.Account),
	}
}

func (i *InMemoryPaymeKeyRepository) FindAccountByID(accountID string) (*domain.Account, error) {
	for _, account := range i.accounts {
		if account.ID == accountID {
			return account, nil
		}
	}

	return nil, domain.ErrAccountNotFound
}

func (i *InMemoryPaymeKeyRepository) SaveAccount(account *domain.Account) error {
	i.Lock()
	defer i.Unlock()
	i.accounts[account.ID] = account
	return nil
}

func (i *InMemoryPaymeKeyRepository) FindKey(key string) (*domain.PaymeKey, error) {
	for _, paymeKey := range i.paymeKeys {
		if paymeKey.Key == key {
			return paymeKey, nil
		}
	}

	return nil, domain.ErrPaymeKeyNotFound
}

func (i *InMemoryPaymeKeyRepository) SavePaymeKey(paymeKey *domain.PaymeKey) error {
	i.Lock()
	defer i.Unlock()
	i.paymeKeys[paymeKey.ID] = paymeKey
	return nil
}
