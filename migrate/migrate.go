package main

import (
	configs "gst.inventario/Backend/conf"
	models "gst.inventario/Backend/model"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Producto{})

}
