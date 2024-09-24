package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	ginSwagger "github.com/swaggo/gin-swagger"

	middleware "mit-api/api/middlewares"
	docs "mit-api/docs"

	swaggerfiles "github.com/swaggo/files"
)

func RegisterRoutes() http.Handler {
	router := gin.Default()

	// add swagger docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	v1 := router.Group("/api/v1")
	v1.Use(middleware.LoggerMiddleware())
	AuthRoutes(v1)
	v1.Use(middleware.JwtAuthMiddleware("my_secret_key"))
	TourRoutes(v1)
	SpinWheelRoutes(v1)

	return router
}
