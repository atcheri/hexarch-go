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

type AppControllersDependencies struct {
	ProjectsRepos    ports.ProjectsRepository
	TranslationsRepo ports.TranslationsRepository
	CommentsRepo     ports.CommentsRepository
}

func NewAppControllers(deps AppControllersDependencies) appControllers {
	return appControllers{
		projectsController:  routes.NewProjectsController(deps.ProjectsRepos),
		languagesController: routes.NewTranslationsController(deps.TranslationsRepo),
		commentsController:  routes.NewCommentsController(deps.CommentsRepo),
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
	routes.AddProjectsRoutes(app, controllers.projectsController)
	routes.AddCommentsRoutes(app, controllers.commentsController)

	app.StaticFS("/api/docs", http.FS(docs.Swagger))

	return app
}
