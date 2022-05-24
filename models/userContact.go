package models

type UserContact struct {
	Id     uint   `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Phone  string `json:"phone" validate:"required,number"`
	UserId uint   `json:"user_id" validate:"required"`
}
