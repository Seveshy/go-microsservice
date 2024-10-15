package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Daniel Major", "daniel@mail.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Daniel Major", client.Name)
	assert.Equal(t, "daniel@mail.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Daniel Major", "daniel@mail.com")
	err := client.Update("Daniel Major Update", "daniel@mail.com")
	assert.Nil(t, err)
	assert.Equal(t, "Daniel Major Update", client.Name)
	assert.Equal(t, "daniel@mail.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Daniel Major", "daniel@mail.com")
	err := client.Update("", "daniel@mail.com")
	assert.Error(t, err, "name is required")
}
