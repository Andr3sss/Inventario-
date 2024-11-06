package controller

import (
	"github.com/gin-gonic/gin"
	configs "gst.inventario/Backend/conf"
	models "gst.inventario/Backend/model"
)

type ProductRequestBody struct {
	Nombre string `json:"nombre"`
	Autor  string `json:"autor"`
	Genero string `json:"genero"`
}

func ProductCreate(c *gin.Context) {

	body := ProductRequestBody{}

	c.BindJSON(&body)

	products := &models.Producto{Nombre: body.Nombre, Autor: body.Autor, Genero: body.Genero}

	result := configs.DB.Create(&products)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Error al insertar producto"})
		return
	}

	c.JSON(200, &products)
}

func ProductGet(c *gin.Context) {
	var products []models.Producto
	configs.DB.Find(&products)
	c.JSON(200, &products)
}

func ProductGetById(c *gin.Context) {
	id := c.Param("id")
	var products models.Producto
	configs.DB.First(&products, id)
	c.JSON(200, &products)
}

func ProductUpdate(c *gin.Context) {

	id := c.Param("id")
	var products models.Producto
	configs.DB.First(&products, id)

	body := ProductRequestBody{}
	c.BindJSON(&body)
	data := &models.Producto{Nombre: body.Nombre, Autor: body.Autor, Genero: body.Genero}

	result := configs.DB.Model(&products).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Fallo al actulizar product"})
		return
	}

	c.JSON(200, &products)
}

func ProductDelete(c *gin.Context) {
	id := c.Param("id")
	var products models.Producto
	configs.DB.Delete(&products, id)
	c.JSON(200, gin.H{"deleted": true})
}
