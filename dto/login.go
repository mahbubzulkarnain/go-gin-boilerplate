package dto

import "time"

// LoginRequest ...
type LoginRequest struct {
	User     string
	Password string
}

// LoginResponse ...
type LoginResponse struct {
	UserID string
	Token  string
}

// JWTConfig ...
type JWTConfig struct {
	Key      string
	Duration time.Duration
}
