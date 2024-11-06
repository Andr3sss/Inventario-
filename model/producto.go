package models

import (
	"gorm.io/gorm"
)

type Producto struct {
	gorm.Model
	ID     uint
	Nombre string
	Autor  string
	Genero string
	Year   int
}

func (Producto) TableName() string {
	return "Producto"
}
