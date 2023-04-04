package routes

import (
	"log"
	"net/http"

	"github.com/atomiccoconut/shortenough/internal/database"
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

	short, err := database.CreateShort(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An internal error occured!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short": short})
}

func ServeShortRoute(c *gin.Context) {
	short, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "I don't know what to say."})
		return
	}

	c.Redirect(http.StatusMovedPermanently, short)
}
