package model_test

import (
	"codepix/domain/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateANewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "4586"
	ownerName := "Juliano Silveira da Costa"
	account, err := model.NewAccount(bank, ownerName, accountNumber)

	require.Nil(t, err)
	require.NotEmpty(t, account.Id)

	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankId, bank.Id)

	_, err = model.NewAccount(bank, "Aaa", ownerName)
	require.Nil(t, err)
}
