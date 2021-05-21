package models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	Message string `json:"message"`
}