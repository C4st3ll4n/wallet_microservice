package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Any name", "anyemail@mail.com")
	account := NewAccount(client)

	assert.NotNil(t, client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestAccount_Credit(t *testing.T) {
	client, _ := NewClient("Any name", "anyemail@mail.com")
	account := NewAccount(client)

	assert.NotNil(t, client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)

	account.Credit(150.0)
	assert.Equal(t, account.Balance, 150.0)
}

func TestAccount_Debit(t *testing.T) {
	client, _ := NewClient("Any name", "anyemail@mail.com")
	account := NewAccount(client)

	assert.NotNil(t, client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)

	account.Credit(150.0)
	assert.Equal(t, account.Balance, 150.0)

	account.Debit(50.0)
	assert.Equal(t, account.Balance, 100.0)
}
