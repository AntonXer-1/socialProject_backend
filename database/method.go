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

func CheckUser(conn *pgx.Conn, ctx context.Context, user_email string, user_role string) (bool, error) {
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
