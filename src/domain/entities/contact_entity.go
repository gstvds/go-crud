package entities

import (
	"time"
)

type CHANNEL string

const (
	EMAIL             CHANNEL = "EMAIL"
	PUSH_NOTIFICATION CHANNEL = "PUSH_NOTIFICATION"
)

type Contact struct {
	Id        string    `json:"id,omitempty"`
	Channel   CHANNEL   `json:"channel,omitempty"`
	Enabled   bool      `json:"enabled,omitempty"`
	Receiver  string    `json:"receiver,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewContact return an instance of Contact
func NewContact(channel CHANNEL, enabled bool, receiver string, userId string) *Contact {
	return &Contact{
		Channel:  channel,
		Enabled:  enabled,
		Receiver: receiver,
		UserId:   userId,
	}
}
