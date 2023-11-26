package usecase

import "github.com/m3k3r1/payme/internal/domain"

type PaymeUseCase struct {
	PaymeKeyRepository domain.PaymeKeyRepository
}

func (p *PaymeUseCase) FindKey(key string) (*domain.PaymeKey, error) {
	paymeKey, err := p.PaymeKeyRepository.FindKey(key)
	if err != nil {
		return nil, err
	}
	return paymeKey, nil
}

func (p *PaymeUseCase) RegisterKey(key string, accountId string) (*domain.PaymeKey, error) {
	account, err := p.PaymeKeyRepository.FindAccountByID(accountId)
	if err != nil {
		return nil, err
	}

	paymeKey, err := domain.NewPaymeKey(key, account)
	if err != nil {
		return nil, err
	}

	err = p.PaymeKeyRepository.SavePaymeKey(paymeKey)
	if err != nil {
		return nil, err
	}

	return paymeKey, nil
}
