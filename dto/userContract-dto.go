package dto

type UserContactDto struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required,number"`
}
