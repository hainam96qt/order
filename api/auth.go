package api

import (
	"encoding/json"
	"net/http"
	"order-gokomodo/middleware"
	"order-gokomodo/pkg/entities"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (a *APIv1) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input LoginRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	result, err := a.AuthService.Login(ctx, &entities.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	middleware.WriteResponse(ctx, w, r, &LoginResponse{Token: result.Token})
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Address  string `json:"address"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func (a *APIv1) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusBadRequest)
		return
	}

	result, err := a.AuthService.Register(ctx, &entities.RegisterRequest{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
		Role:     input.Role,
		Address:  input.Address,
	})
	if err != nil {
		middleware.WriteError(w, r, err, http.StatusInternalServerError)
		return
	}
	middleware.WriteResponse(ctx, w, r, &RegisterResponse{Token: result.Token})
}
