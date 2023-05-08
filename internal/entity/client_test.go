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
