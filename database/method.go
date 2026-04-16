package database

import (
	"context"
	"log"
	"socialProject_backend/user"

	"github.com/jackc/pgx/v5"
)

func AddUser(conn *pgx.Conn, ctx context.Context, new_user *user.User) error {
	query := `
			INSERT INTO users (email, password, full_name, photo_url, role) 
			VALUES ($1, $2, $3, $4, $5)
			`

	_, err := conn.Exec(ctx, query,
		new_user.Email,
		new_user.Password,
		new_user.FullName,
		new_user.PhotoURL,
		new_user.Role)

	if err != nil {
		log.Printf("Error adding user: %v", err)
		return err
	}

	return nil
}

func CheckRegisterUser(conn *pgx.Conn, ctx context.Context, user_email, user_role string) (bool, error) {
	query := `
		SELECT EXISTS (
		    SELECT 1
			FROM users
			WHERE email = $1 AND role = $2)
`

	var exists bool

	err := conn.QueryRow(ctx, query, user_email, user_role).Scan(&exists)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return false, err
	}

	return exists, nil
}

func CheckUserAuthorization(conn *pgx.Conn, ctx context.Context, user_email, user_role string) (*user.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE email = $1 AND role = $2
		`

	var user user.User
	err := conn.QueryRow(ctx, query, user_email, user_role).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.PhotoURL,
		&user.Role,
		&user.CreatedAt)

	if err != nil {
		log.Printf("Error checking user: %v", err)
		return nil, err
	}

	return &user, nil
}
