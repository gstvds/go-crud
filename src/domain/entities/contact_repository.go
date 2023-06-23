package entities

import (
	"context"
)

type ContactRepository interface {
	Create(*Contact, context.Context) (*Contact, error)
	GetByUserId(string, context.Context) (*[]Contact, error)
}
