package internaluser

import (
	"context"
	"database/sql"
	"fmt"
	"ws_postgre/helper"
)

type UserRepositoryIPLM struct{}

func NewUserRepositoryIPLM() *UserRepositoryIPLM {
	return &UserRepositoryIPLM{}
}

func (repo *UserRepositoryIPLM) Save(ctx context.Context, tx *sql.Tx, req User) User {
	query := "INSERT INTO users (id, name, email, password, created_at) values($1,$2,$3,$4,$5)"
	fmt.Println("Actual SQL Query:", query)
	_, err := tx.ExecContext(ctx, query, req.ID, req.Name, req.Email, req.Password, req.CreatedAt)

	helper.HelperError(err, "error creating user repository")
	fmt.Print("userr gblokk", err)

	return req

}
func (repo *UserRepositoryIPLM) GetByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (User, error) {
	row := tx.QueryRowContext(ctx, "select id, name, email, created_at from users where email = $1", userEmail)

	user := User{}
	if row.Err() != nil {
		return User{}, row.Err()
	}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	helper.HelperError(err, "error scanning getByEmail user repository")
	return user, nil

}

func (repo *UserRepositoryIPLM) GetByID(ctx context.Context, tx *sql.Tx, userID string) (User, error) {
	row := tx.QueryRowContext(ctx, "select id, name, email, created_at from users where id = $1", userID)

	user := User{}
	if row.Err() != nil {
		return User{}, row.Err()
	}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	helper.HelperError(err, "error scanning getByID user repository")
	return user, nil

}
func (repo *UserRepositoryIPLM) GetAll(ctx context.Context, tx *sql.Tx) []User {
	rows, err := tx.QueryContext(ctx, "select id, name, email, created_at from users")
	helper.HelperError(err, "error querying getAll user repository")

	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		helper.HelperError(err, "error scanning getAll user repository")
		users = append(users, user)
	}

	return users

}
