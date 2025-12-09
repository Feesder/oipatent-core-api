package dto

import "time"

type TokensDto struct {
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	AccessExpires  time.Time `json:"access_expires"`
	RefreshExpires time.Time `json:"refresh_expires"`
}

type TokenResponse struct {
	AccessToken   string    `json:"access_token"`
	AccessExpires time.Time `json:"access_expires"`
	TokenType     string    `json:"token_type"`
}
