package dto

import "strings"

type UserDTO struct {
	fullName string
	eMail    string
	password string
	phone    string
	photoURL string
	role     string
}

// ниже - функции для валидации корректности ввёденных пользователем данных
// для регистрации и авторизации соответственно
func (u *UserDTO) ValidationForRegistration() error {
	if u.fullName == "" {
		return ErrFullNameIsEmpty
	}
	if u.eMail == "" || strings.Contains(u.eMail, "@gmail.com") {
		return ErrInvalidEmail
	}
	if u.password == "" {
		return ErrPasswordIsEmpty
	}
	if u.role == "" {
		return ErrRoleIsEmpty
	}
	return nil
}

func (u *UserDTO) ValidationForAuthorization() error {
	if u.eMail == "" || strings.Contains(u.eMail, "@gmail.com") {
		return ErrInvalidEmail
	}
	if u.password == "" {
		return ErrPasswordIsEmpty
	}
	return nil
}
