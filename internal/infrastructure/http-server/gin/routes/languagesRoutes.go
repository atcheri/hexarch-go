package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	dto "github.com/atcheri/hexarch-go/internal/core/dtos"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
)

// TranslationsController is the controller for the translations route
type TranslationsController struct {
	translationsRepo ports.TranslationsRepository
}

// AddTranslationsRoutes adds the routes to the translations endpoint
func AddTranslationsRoutes(router *gin.Engine, tc TranslationsController) {
	group := router.Group("/translations/:projectName")
	group.GET("/", tc.GetAllHandler)
}

// NewTranslationsController is a TranslationsController factory function
func NewTranslationsController(repo ports.TranslationsRepository) TranslationsController {
	return TranslationsController{
		translationsRepo: repo,
	}
}

// GetAllHandler is the http request handler that returns all translations for a given project
func (lc TranslationsController) GetAllHandler(c *gin.Context) {
	name := c.Param("projectName")
	offset := 0
	limit := 5
	translations, err := lc.translationsRepo.GetForProject(c, name, offset, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error_type": http.StatusText(http.StatusNotFound),
			"name":       "Resource not found",
			"message":    fmt.Sprintf("the translations were not found for this project: %s", name),
		})
		return
	}
	dots := dto.ToTranslationKeyDTOs(translations)
	c.JSON(http.StatusOK, gin.H{
		"translations": dots,
		// TODO: we need to calculate the total translations for the project
		"total": len(dots),
	})
}
