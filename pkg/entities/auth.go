package entities

import "context"

type Authentication interface {
	Login (context.Context, *LoginRequest) (*LoginResponse, error)
	Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error)
}

type LoginRequest struct {
	Email string
	Password string
}

type LoginResponse struct {
	Token string
}

type RegisterRequest struct {
	Email string
	Password string
	Name string
	Role string
	Address string
}

type RegisterResponse struct {
	Token string
}
