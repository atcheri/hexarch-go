package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/atcheri/hexarch-go/docs"
)

//// GinApp will implement the Server interface
//type GinApp struct{}

func NewGinApp() *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	app.StaticFS("/api/docs", http.FS(docs.Swagger))

	return app
}
