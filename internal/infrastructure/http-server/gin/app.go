package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/atcheri/hexarch-go/docs"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/atcheri/hexarch-go/internal/infrastructure/http-server/gin/routes"
)

type appControllers struct {
	languagesController routes.TranslationsController
}

type AppControllersDependencies struct {
	TranslationsRepo ports.TranslationsRepository
}

func NewAppControllers(deps AppControllersDependencies) appControllers {
	return appControllers{
		languagesController: routes.NewTranslationsController(deps.TranslationsRepo),
	}
}

func NewGinApp(controllers appControllers) *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"date": time.Now().String()})
	})

	routes.AddTranslationsRoutes(app, controllers.languagesController)

	app.StaticFS("/api/docs", http.FS(docs.Swagger))

	return app
}
