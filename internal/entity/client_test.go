package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john@email.com", client.Email)

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@email.com")
	err = client.Update("John Doe Updated", "new.john@emai.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Updated", client.Name)
	assert.Equal(t, "new.john@emai.com", client.Email)
}

func TestUpdateClientWhenNewArgsAreInvalid(t *testing.T) {
	client, err := NewClient("John Doe", "john@email.com")
	err = client.Update("", "")
	assert.Error(t, err, "name is required")
	assert.NotNil(t, err)
}

func TestAddAccountToAClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@email.com")
	account := NewAccount(client)

	client.AddAccount(account)
	assert.NotNil(t, client)
	assert.NotNil(t, account)
	assert.Equal(t, 1, len(client.Accounts))
	assert.Equal(t, account, client.Accounts[0])
}

func TestAddAccountToAClientWhenAccountIsFromAnotherClient(t *testing.T) {
	client_1, _ := NewClient("John Doe", "john@email.com")
	client_2, _ := NewClient("Josh", "josh@email.com")
	account := NewAccount(client_1)

	err := client_2.AddAccount(account)
	assert.Equal(t, err.Error(), "account does not belong to this client")
	assert.Nil(t, client_2.Accounts)
}
