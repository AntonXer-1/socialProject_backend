package user

import (
	"time"
)

type User struct {
	FullName string
	Email    string
	Password string
	Phone    string
	PhotoURL string
	Role     string

	CreatedAt time.Time
}

func NewUser(
	fullName, email, password, phone, photoURL, role string) *User {
	return &User{
		FullName: fullName,
		Email:    email,
		Password: password,
		Phone:    phone,
		PhotoURL: photoURL,
		Role:     role,

		CreatedAt: time.Now(),
	}
}
