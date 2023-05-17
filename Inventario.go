package main

import "time"

type Inventory struct {
	ID          uint      `json:"id" gorm:"column:Id_Inv"`
	Description string    `json:"description" gorm:"column:Descripcion"`
	Date        time.Time `json:"date" gorm:"column:Fecha"`
	Planta      string    `json:"planta" gorm:"column:Planta"`
	Bodega      string    `json:"bodega" gorm:"column:bodega"`
	User        uint      `json:"userId" gorm:"column:id_usuario"`
	Canceled    int       `json:"canceled" gorm:"column:Anulado"`
}

func (Inventory) TableName() string {
	return "inventario_enc"
}
