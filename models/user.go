package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           uint          `json:"id"`
	UserName     string        `json:"username"`
	Email        string        `json:"email" gorm:"unique"`
	Token        string        `json:"token"`
	Password     []byte        `json:"-"`
	UserContacts []UserContact `json:"user_contact" gorm:"foreignKey:UserId"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
