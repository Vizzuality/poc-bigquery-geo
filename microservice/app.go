package microservice

import (
	"os"

	"github.com/gin-gonic/gin"
)

// InitApp ...
func InitApp() {

	router := gin.Default()
	router.GET("/query", queryRouter)
	router.StaticFile("/", "./microservice/assets/index.html")
	router.Static("/assets", "microservice/assets")

	// in prod using $PORT in dev using 3001 to map gin $PORT:3001
	port := os.Getenv("PORT")
	if os.Getenv("GIN_MODE") == "debug" {
		port = "3001"
	}
	router.Run(":" + port)

}
