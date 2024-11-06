package rutas

import (
	"github.com/gin-gonic/gin"
	controllers "gst.inventario/Backend/controller"
)

func ProductRouter(router *gin.Engine) {

	routes := router.Group("api/v1/productos")
	routes.POST("", controllers.ProductCreate)
	routes.GET("", controllers.ProductGet)
	routes.GET("/:id", controllers.ProductGetById)
	routes.PUT("/:id", controllers.ProductUpdate)
	routes.DELETE("/:id", controllers.ProductDelete)

}
