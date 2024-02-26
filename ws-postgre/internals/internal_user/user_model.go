package internaluser

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type CreateUser struct {
	Name     string `js:"name"`
	Email    string `js:"email"`
	Password string `js:"password"`
}

type UserResponse struct {
	ID        string    `js:"id"`
	Name      string    `js:"name"`
	Email     string    `js:"email"`
	CreatedAt time.Time `js:"created_at"`
}

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}
