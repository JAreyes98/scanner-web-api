package main

import "time"

type Order struct {
	ID         string    `json:"id" gorm:"column:PEDIDO"`
	Status     string    `json:"status" gorm:"column:ESTADO"`
	OrderDate  time.Time `json:"orderDate" gorm:"column:FECHA_PEDIDO"`
	Route      string    `json:"route" gorm:"column:RUTA"`
	SellerCode string    `json:"sellerCode" gorm:"column:VENDEDOR"`
	Customer   string    `json:"customer" gorm:"column:CLIENTE"`
}

func (Order) TableName() string {
	return "pedido"
}
