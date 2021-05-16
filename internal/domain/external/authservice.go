package external

import "anest.top/youli/internal/data/models"

type authService interface {
	PasswordLogin(account, password string) (*models.Auth, error)
}
