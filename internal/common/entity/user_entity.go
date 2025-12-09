package entity

import "time"

type User struct {
	Id        string    `db:"id"`
	Firstname string    `db:"firstname"`
	Lastname  string    `db:"lastname"`
	Role      string    `db:"role"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
