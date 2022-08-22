package models

import "time"

type User struct {
	Email     string `gorm:"type:varchar(50);primaryKey" json:"email,omitempty"`
	Name      string `gorm:"type:varchar(50);not null" json:"name,omitempty"`
	BirthDate string `gorm:"type:varchar(20);not null" json:"birth_date,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
