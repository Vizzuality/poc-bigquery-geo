package microservice

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// InitApp ...
func InitApp() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "123")
	})
	router.GET("/user/:name", getUserRouter)

	// in prod using $PORT in dev using 3001 to map gin $PORT:3001
	port := os.Getenv("PORT")
	if os.Getenv("GIN_MODE") == "debug" {
		port = "3001"
	}
	router.Run(":" + port)

}
