package routers

import (
	"mit-api/internal/api/middlewares"
	"net/http"

	docs "mit-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes() http.Handler {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	// add swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))
	AuthRoutes(v1)
	v1.Use(middlewares.JwtAuthMiddleware("my_secret_key"))
	TourRoutes(v1)
	SpinWheelRoute(v1)

	return router
}
