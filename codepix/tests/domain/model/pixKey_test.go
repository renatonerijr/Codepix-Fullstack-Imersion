package model_test

import (
	"codepix/domain/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateANewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "4586"
	ownerName := "Juliano Silveira da Costa"
	account, err := model.NewAccount(bank, ownerName, accountNumber)

	kind := model.EMAIL
	key := "jsilveira@gmail.com"
	pixKey, err := model.NewPixKey(kind, account, key)

	require.Nil(t, err)
	require.NotEmpty(t, pixKey.Id)
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, model.ACTIVE)

	kind = model.CPF
	_, err = model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
