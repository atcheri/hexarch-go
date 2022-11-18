// Package server wires up the Gin framework dependencies
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
	projectsController  routes.ProjectsController
	languagesController routes.TranslationsController
	commentsController  routes.CommentsController
}

// AppControllersDependencies lists up the controller dependencies, mostly repositories
type AppControllersDependencies struct {
	ProjectsRepos    ports.ProjectsRepository
	TranslationsRepo ports.TranslationsRepository
	CommentsRepo     ports.CommentsRepository
}

// NewAppControllers is the AppController factory function
func NewAppControllers(deps AppControllersDependencies) appControllers {
	return appControllers{
		projectsController:  routes.NewProjectsController(deps.ProjectsRepos),
		languagesController: routes.NewTranslationsController(deps.TranslationsRepo),
		commentsController:  routes.NewCommentsController(deps.CommentsRepo),
	}
}

// NewGinApp sets up the Gin engine
func NewGinApp(controllers appControllers) *gin.Engine {
	app := gin.Default()

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"date": time.Now().String()})
	})

	apiGroup := app.Group("/api")
	routes.AddTranslationsRoutes(apiGroup, controllers.languagesController)
	routes.AddProjectsRoutes(apiGroup, controllers.projectsController)
	routes.AddCommentsRoutes(apiGroup, controllers.commentsController)

	app.StaticFS("/api/docs", http.FS(docs.Swagger))

	return app
}
