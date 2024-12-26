package models

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Password  string    `gorm:"not null" json:"password"` // Omit when sending over the wire
	Email     string    `gorm:"not null;unique" json:"email"`
	Role      string    `gorm:"type:varchar(255);not null;default:'user'" json:"role"` // e.g., "super_admin", "user"
	DiscordID *string   `gorm:"type:uuid;default:null" json:"discord_id,omitempty"`    // Optional foreign key
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
