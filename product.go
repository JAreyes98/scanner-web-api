package main

import "time"

type Product struct {
	Id          int    `json:"id" gorm:"column:ID"`
	Name        string `json:"name" gorm:"column:NOMBRE"`
	Barcode     string `json:"barcode" gorm:"column:CodBarraSupermercado"`
	CodProducto string `json:"codproduct" gorm:"column:CODPRODUCTO"`
	//Stock    int       `json:"stock"`
	CreateAt time.Time `json:"createAt" gorm:"column:Fecha_Crea"`
	UpdateAt time.Time `json:"updateAt" gorm:"column:Fecha_Modifica"`
}

type ProductCreate struct {
}
