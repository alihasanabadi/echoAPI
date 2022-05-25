package models

type Blog struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type Users struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Token    string `json:"token"`
}
