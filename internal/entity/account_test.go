package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, err := NewClient("John Doe", "john@email.com")
	if err != nil {
		err.Error()
	}
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {

	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@email.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}


func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@email.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}