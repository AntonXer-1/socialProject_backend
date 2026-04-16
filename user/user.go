package user

import (
	"time"
)

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	PhotoURL string `json:"photoURL"`
	Role     string `json:"role"`

	CreatedAt time.Time `json:"created_at"`
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
