package user

import (
	"time"
)

type User struct {
	fullName string
	eMail    string
	password string
	phone    string
	photoURL string
	role     string

	createdAt time.Time
}

func NewUser(
	fullName, eMail, password, phone, photoURL, role string) *User {
	return &User{
		fullName: fullName,
		eMail:    eMail,
		password: password,
		phone:    phone,
		photoURL: photoURL,
		role:     role,

		createdAt: time.Now(),
	}
}
