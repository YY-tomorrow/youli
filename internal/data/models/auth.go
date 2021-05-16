package models

import "gorm.io/gorm"

type Auth struct {
	Account  string `json:"account,omitempty" bson:"account"`
	Password string `json:"password,omitempty" bson:"password"`
	NickName string `json:"nick_name,omitempty" bson:"nick_name"`
	gorm.Model
}

func (t *Auth) TableName() string {
	return "auth"
}
