package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Jonh Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Jonh Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err) // expect error
	assert.Nil(t, client) // not expect client
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Jonh Doe", "j@j.com")
	err := client.Update("Jonh Doe Update", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jonh Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)

}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Jonh Doe", "j@j.com")
	err := client.Update("", "j@j.com")
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Jonh Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
