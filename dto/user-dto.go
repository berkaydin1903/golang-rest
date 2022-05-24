package dto

type UserRegisterDto struct {
	UserName        string `json:"username" validate:"required,alpha" `
	Email           string `json:"email" gorm:"unique" validate:"required,email"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
	Password        string `json:"password" validate:"required"`
}
