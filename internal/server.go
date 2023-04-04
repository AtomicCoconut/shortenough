package internal

import (
	"github.com/atomiccoconut/shortenough/internal/routes"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.GET("/:id", routes.ServeShort)
	router.POST("/", routes.CreateShort)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
