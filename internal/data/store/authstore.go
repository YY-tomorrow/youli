package store

import (
	"anest.top/youli/internal/data/models"
	"gorm.io/gorm"
)

type AuthStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{db: db}
}

func (t *AuthStore) PasswordLogin(account, password string) (*models.Auth, error) {
	a := &models.Auth{}
	return a, t.db.Find(a, "account = ? and password = ?", account, password).Error
}
