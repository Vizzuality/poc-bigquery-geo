package microservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryRouter(c *gin.Context) {
	log.Println("[ROUTER] Query")
	sql := c.Query("sql")
	result, err := queryService(sql)

	if err != nil {
		c.JSON(500, "Generic Error")
	} else {
		c.JSON(http.StatusOK, result)
	}
}
