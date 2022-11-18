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
func AddProjectsRoutes(apiGroup *gin.RouterGroup, tc ProjectsController) {
	group := apiGroup.Group("/projects")
	group.POST("/", tc.PostProjectHandler)
	group.PUT("/:projectName", tc.PutProjectHandler)
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
			"impossible to create the new project",
		))
		return
	}

	if err := lc.projectsRepo.Create(c, body.Name); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			fmt.Sprintf("impossible to create a new project called %s", body.Name),
		))
		return
	}

	c.Status(http.StatusCreated)
}

func (lc ProjectsController) PutProjectHandler(c *gin.Context) {
	oldName := c.Param("projectName")
	var body dto.EditProjectRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not updated",
			"impossible to edit the project ",
		))
		return
	}

	newName := body.Name
	if err := lc.projectsRepo.Edit(c, oldName, newName); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not updated",
			fmt.Sprintf("impossible to edit the new project called %s to %s", oldName, newName),
		))
		return
	}

	c.Status(http.StatusNoContent)
}
