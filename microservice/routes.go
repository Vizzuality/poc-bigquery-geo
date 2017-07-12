package microservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserRouter(c *gin.Context) {
	log.Println("[ROUTER] Getting user")
	name := c.Param("name")
	user := getUserService(name)
	c.JSON(http.StatusOK, user)
}
