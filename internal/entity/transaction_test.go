package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Any client", "ac@mail.com")
	client2, _ := NewClient("Other Client", "oc@mail.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1500.0)
	account2.Credit(500.0)

	transaction, err := NewTransaction(account2, account1, 250.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)

	assert.Equal(t, 1750.00, account1.Balance)
	assert.Equal(t, 250.0, account2.Balance)
}

func TestCreateTransactionWithZeroFunds(t *testing.T) {
	client1, _ := NewClient("Any client", "ac@mail.com")
	client2, _ := NewClient("Other Client", "oc@mail.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1500.0)

	transaction, err := NewTransaction(account2, account1, 250.0)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)

	assert.Equal(t, 1500.0, account1.Balance)
	assert.Equal(t, 0.0, account2.Balance)
}
