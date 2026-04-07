package dto

import "strings"

type UserDTO struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
	Phone    string `json:"phone"`
	PhotoURL string `json:"photo"`
	Role     string `json:"role"`
}

// ниже - функции для валидации корректности ввёденных пользователем данных
// для регистрации и авторизации соответственно
func (u *UserDTO) ValidationForRegistration() error {
	if u.FullName == "" {
		return ErrFullNameIsEmpty
	}
	if u.Email == "" || !strings.Contains(u.Email, "@gmail.com") {
		return ErrInvalidEmail
	}
	if u.Password == "" {
		return ErrPasswordIsEmpty
	}
	if u.Role == "" {
		return ErrRoleIsEmpty
	}
	return nil
}

func (u *UserDTO) ValidationForAuthorization() error {
	if u.Email == "" || strings.Contains(u.Email, "@gmail.com") {
		return ErrInvalidEmail
	}
	if u.Password == "" {
		return ErrPasswordIsEmpty
	}
	return nil
}

type UserResponseDTO struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
