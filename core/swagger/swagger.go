package swagger

import (
	"simple-conf/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(router *gin.RouterGroup) {
	docs.SwaggerInfo.Title = "Simple Conf Project"
	docs.SwaggerInfo.Description = "Sample Video Conf project by gin framework. Clear structure, custom command, gorm orm, socketIO, swagger, scheduler"
	docs.SwaggerInfo.Version = "0.1.0"

	router.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(-1),
		),
	)
}
