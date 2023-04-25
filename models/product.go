package models

import "time"

type Product struct {
	Id       int       `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Stock    int       `json:"stock"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `json:"update_at" gorm:"autoUpdateTime"`
}

type CreateProduct struct {
	Name  string
	Stock int
}

type AddStock struct {
	Id     int
	Amount int
}
