package main

type Bodega struct {
	Codigo string `json:"codigo" gorm:"column:BODEGA"`
	Nombre string `json:"nombre" gorm:"column:NOMBRE"`
}

func (Bodega) TableName() string {
	return "bodega"
}
