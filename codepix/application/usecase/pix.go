package usecase

import (
	"codepix/domain/model"
	"errors"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.Register(pixKey)
	if pixKey.Id == "" {
		return nil, errors.New("unable to create new pix key")
	}
	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
