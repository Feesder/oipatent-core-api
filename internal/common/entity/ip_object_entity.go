package entity

import "time"

type IpObject struct {
	Id           string    `db:"id"`
	UserId       string    `db:"user_id"`
	Title        string    `db:"title"`
	PatentType   string    `db:"patent_type"`
	Description  string    `db:"description"`
	Jurisdiction string    `db:"jurisdiction"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
