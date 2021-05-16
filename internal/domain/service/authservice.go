package service

import "anest.top/youli/internal/domain/external"

type AuthService struct {
	Register *external.Register
}

func NewAuthService(register *external.Register) *AuthService {
	return &AuthService{Register: register}
}
