package routes

import (
	"fmt"
	"net/http"

	dto "github.com/atcheri/hexarch-go/internal/core/dtos"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/gin-gonic/gin"
)

// ProjectsController is the controller for the projects route
type ProjectsController struct {
	projectsRepo ports.ProjectsRepository
}

// AddProjectsRoutes adds the routes to the projects endpoint
func AddProjectsRoutes(router *gin.Engine, tc ProjectsController) {
	group := router.Group("/projects")
	group.POST("/", tc.PostProjectHandler)
}

// NewProjectsController is a ProjectsController factory function
func NewProjectsController(repo ports.ProjectsRepository) ProjectsController {
	return ProjectsController{
		projectsRepo: repo,
	}
}

func (lc ProjectsController) PostProjectHandler(c *gin.Context) {
	var body dto.CreateProjectRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			"impossible to create the new project %s",
		))
		return
	}

	if err := lc.projectsRepo.Create(c, body.Name); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			fmt.Sprintf("impossible to create a new the new project called %s", body.Name),
		))
		return
	}

	c.Status(http.StatusCreated)
}
