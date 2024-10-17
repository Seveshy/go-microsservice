package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Daniel", "daniel@mail.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Daniel", "daniel@mail.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Daniel", "daniel@mail.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(100), account.Balance)
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Daniel", "daniel@mail.com")
	account := NewAccount(client)
	err := client.Addcount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Account))
}
