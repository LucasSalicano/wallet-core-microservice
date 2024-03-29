package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID       string
	Name     string
	Email    string
	createAt time.Time
	updateAt time.Time
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		createAt: time.Now(),
		updateAt: time.Now(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.updateAt = time.Now()
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}
