package models

type Camion struct {
	ID         uint   `json:"id" gorm:"column:Camion"`
	Nombre     string `json:"nombre" gorm:"column:NOMBRE"`
	Placa      string `json:"placa" gorm:"column:placa"`
	PesoMaximo int    `json:"pesoMaximo" gorm:"column:Peso_MAX"`
	PesoMinimo int    `json:"pesoMinimo" gorm:"column:Peso_min"`
}

func (Camion) TableName() string {
	return "Camiones"
}
