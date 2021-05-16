package external

import "anest.top/youli/internal/data/models"

type authStore interface {
	PasswordLogin(account, password string) (*models.Auth, error)
}
