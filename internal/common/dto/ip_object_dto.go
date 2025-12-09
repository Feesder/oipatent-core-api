package dto

import "time"

type CreateIpObjectDto struct {
	UserId       string `json:"user_id" validate:"required"`
	Title        string `json:"title" validate:"required,min=3,max=50"`
	Description  string `json:"description" validate:"required,min=10,max=1024"`
	Jurisdiction string `json:"jurisdiction" validate:"required,min=1,max=10"`
	PatentType   string `json:"patent_type" validate:"required,min=1,max=50"`
}

type IpObjectDto struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	Title        string    `json:"title"`
	PatentType   string    `json:"patent_type"`
	Description  string    `json:"description"`
	Jurisdiction string    `json:"jurisdiction"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
