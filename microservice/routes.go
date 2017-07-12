package microservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryRouter(c *gin.Context) {
	log.Println("[ROUTER] Query")
	sql := c.Query("sql")
	result := queryService(sql)
	c.JSON(http.StatusOK, result)
}
