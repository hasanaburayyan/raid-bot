package models

import "time"

type Group struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"not null;unique" json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Members []GroupMember `gorm:"constraint:OnDelete:CASCADE" json:"members,omitempty"`
}
