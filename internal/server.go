package internal

import (
	"net/http"

	"github.com/atomiccoconut/shortenough/internal/database"
	"github.com/atomiccoconut/shortenough/internal/routes"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	database.Connect()

	router := gin.Default()

	router.LoadHTMLFiles("public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/:id", routes.ServeShortRoute)

	apiGroup := router.Group("/api/v1/")
	apiGroup.POST("/create", routes.CreateShortRoute)

	err := router.Run()

	if err != nil {
		panic(err)
	}
}
