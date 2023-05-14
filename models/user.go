package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	//gorm.Model
	ID       uint   `json:"id" gorm:"column:Id_Usuario"`
	Name     string `json:"name" gorm:"column:Nombre"`
	Username string `json:"username" gorm:"column:login"`
	Password string `json:"password" gorm:"column:Contraseña"`
	Baja     uint   `json:"isBaja" gorm:"column:baja"`
}

func (user *Usuario) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *Usuario) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
