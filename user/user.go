package user

import (
	"time"
)

type User struct {
	fullName string
	email    string
	password string
	phone    string
	photoURL string
	role     string

	createdAt time.Time
}

func NewUser(
	fullName, email, password, phone, photoURL, role string) *User {
	return &User{
		fullName: fullName,
		email:    email,
		password: password,
		phone:    phone,
		photoURL: photoURL,
		role:     role,

		createdAt: time.Now(),
	}
}
