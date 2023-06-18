package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;type:varchar(36)" json:"id,omitempty"`
	Email     string    `gorm:"type:varchar(200)" json:"email,omitempty"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name,omitempty"`
	BirthDate string    `gorm:"type:varchar(20);not null" json:"birth_date,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	// generate UUID
	uuid := uuid.New()
	user.Id = uuid
	return nil
}
