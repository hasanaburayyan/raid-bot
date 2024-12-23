package models

type Class struct {
	Name  string   `json:"name"`
	Specs []string `json:"specs"`
	Roles []string `json:"roles"`
}
