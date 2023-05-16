package main

type ProductionInventoryDto struct {
	Date         string `json:"date"`
	Planta       Planta `json:"planta"`
	Barcode      string `json:"barcode"`
	IdInventario string `json:"idInventario"`
}

type ChargeInventory struct {
	Date    string `json:"date"`
	Camion  Camion `json:"camion"`
	OrderId int    `json:"orderId"`
	Barcode string `json:"barcode"`
}

type WarehouseInventory struct {
	Date         string `json:"date"`
	Bodega       Bodega `json:"bodega"`
	Barcode      string `json:"barcode"`
	IdInventario string `json:"idInventario"`
}
