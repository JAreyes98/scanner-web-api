package models

type Planta struct {
	Planta     string `json:"planta" gorm:"column:Planta"`
	Nombre     string `json:"nombre" gorm:"column:NOMBRE"`
	Produccion int    `json:"produccion" gorm:"column:Produccion"`
}

func (Planta) TableName() string {
	return "planta"
}
