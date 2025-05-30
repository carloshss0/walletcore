package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@email.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Mark Doe", "mark@email.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)
}

func TestCreateTransactionWhenAccountFromBalanceIsInsufficient(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@email.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Mark Doe", "mark@email.com")
	account2 := NewAccount(client2)

	account1.Credit(100)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 1000)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Equal(t, 100.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionWhenAmountIsNegative(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@email.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Mark Doe", "mark@email.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, -10)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}
