package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

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
	group.POST("/", tc.PostProjectTranslationHandler)
	group.DELETE("/", tc.DeleteProjectTranslationHandler)

	group.PUT("/:translationId", tc.PutProjectTranslationHandler)
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_type": http.StatusText(http.StatusNotFound),
			"name":       "Resource not found",
			"message":    errors.Wrap(err, fmt.Sprintf("the translations were not found for this project: %s", name)),
		})
	}
	dots := dto.ToTranslationKeyDTOs(translations)
	c.JSON(http.StatusOK, gin.H{
		"translations": dots,
		// TODO: we need to calculate the total translations for the project
		"total": len(dots),
	})
}

func (lc TranslationsController) PostProjectTranslationHandler(c *gin.Context) {
	name := c.Param("projectName")
	var body dto.CreateProjectTranslationRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			fmt.Sprintf("impossible to create a new translation for this project: %s", name),
		))
		return
	}
	key := body.Key
	code := body.Code
	text := body.Text
	err := lc.translationsRepo.AddForProject(c, name, key, code, text)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			fmt.Sprintf("failed to create a new translation for this project: %s", name),
		))
		return
	}

	c.Status(http.StatusCreated)
}

func (lc TranslationsController) PutProjectTranslationHandler(c *gin.Context) {
	name := c.Param("projectName")
	id := c.Param("translationId")
	var body dto.EditProjectTranslationRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not updated",
			fmt.Sprintf("impossible to edit a translation for this project: %s", name),
		))
		return
	}
	key := body.Key
	code := body.Code
	text := body.Text
	err := lc.translationsRepo.EditForProject(c, id, name, key, code, text)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not updated",
			fmt.Sprintf("failed to edit the translation for this project: %s", name),
		))
		return
	}

	c.Status(http.StatusOK)
}

func (lc TranslationsController) DeleteProjectTranslationHandler(c *gin.Context) {
	name := c.Param("projectName")
	var body dto.DeleteProjectTranslationRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not deleted",
			fmt.Sprintf("impossible to delete the translations for this project: %s", name),
		))
		return
	}
	key := body.Key
	err := lc.translationsRepo.DeleteKeyForProject(c, name, key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not updated",
			fmt.Sprintf("failed to delete the translations for this project: %s, key: %s", name, key),
		))
		return
	}

	c.Status(http.StatusNoContent)

}

func buildGinErrorJSON(code int, name, message string) gin.H {
	return gin.H{
		"error_type": http.StatusText(code),
		"name":       name,
		"message":    message,
	}
}
