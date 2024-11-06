package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	configs "gst.inventario/Backend/conf"
	routes "gst.inventario/Backend/rutas"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.ProductRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ola, bienvenido a mi programa :)",
		})
	})
	r.Run()
}
