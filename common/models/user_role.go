package models

type UserRole struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID string `gorm:"not null;index" json:"user_id"`
	RoleID string `gorm:"not null;index" json:"role_id"`

	// Foreign keys
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE" json:"-"`
}
