package models

type Class struct {
	Name  string   `gorm:"type:text" json:"name"`
	Specs []string `gorm:"type:text[]" json:"specs"`
	Roles []string `gorm:"type:text[]" json:"roles"`
}
