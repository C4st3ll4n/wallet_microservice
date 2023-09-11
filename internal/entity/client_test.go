package entity_test

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := entity.NewClient("Any name", "anyemail@mail.com")
	assert.Nil(t, err)

	assert.Equal(t, "Any name", client.Name)
	assert.Equal(t, "anyemail@mail.com", client.Email)

}

func TestNewClientWithInvalidArgs(t *testing.T) {
	client, err := entity.NewClient("", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestClient_Update(t *testing.T) {
	client, _ := entity.NewClient("Any name", "anyemail@mail.com")
	err := client.Update("Updated", "updated@mail.com")

	assert.Nil(t, err)
	assert.Equal(t, "Updated", client.Name)
	assert.Equal(t, "updated@mail.com", client.Email)
}

func TestClient_UpdateWithInvalidArgs(t *testing.T) {
	client, _ := entity.NewClient("Any name", "anyemail@mail.com")
	err := client.Update("", "")

	assert.NotNil(t, err)

}

func TestClient_AddAccount(t *testing.T) {
	client, _ := entity.NewClient("Any name", "anyemail@mail.com")
	account := entity.NewAccount(client)

	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, len(client.Accounts), 1)
}
