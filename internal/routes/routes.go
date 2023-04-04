package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortRoute(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please provide an url as a query parameter!",
		})
		return
	}
}

func ServeShortRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Serve"})

}
