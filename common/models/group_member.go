package models

type GroupMember struct {
	ID      string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	GroupID string `gorm:"not null" json:"group_id"`
	UserID  string `gorm:"not null" json:"user_id"`
	IsAdmin bool   `gorm:"default:false" json:"is_admin"`

	// Relationships
	Group Group `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group,omitempty"`
	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}
