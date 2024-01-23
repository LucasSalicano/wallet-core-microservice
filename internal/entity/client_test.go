package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "jd@.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, client.ID)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "jd@.com", client.Email)
}

func TestNewClientEmptyEmail(t *testing.T) {
	_, err := NewClient("John Doe", "")
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}

func TestNewClientEmptyName(t *testing.T) {
	_, err := NewClient("", "jd@.com")
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@.com")
	err := client.Update("John Doe 2", "jd2@.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe 2", client.Name)
	assert.Equal(t, "jd2@.com", client.Email)
}

func TestUpdateClientEmptyEmail(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@.com")
	err := client.Update("John Doe 2", "")
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}

func TestUpdateClientEmptyName(t *testing.T) {
	client, _ := NewClient("John Doe", "jd@.com")
	err := client.Update("", "jd2@.com")
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}
