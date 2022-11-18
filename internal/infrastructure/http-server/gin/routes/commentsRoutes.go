package routes

import (
	"fmt"
	"net/http"

	"github.com/atcheri/hexarch-go/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	dto "github.com/atcheri/hexarch-go/internal/core/dtos"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
)

// CommentsController is the controller for the projects route
type CommentsController struct {
	commentsRepo     ports.CommentsRepository
	translationdRepo ports.TranslationsRepository
}

// AddCommentsRoutes adds the routes to the comments endpoint
func AddCommentsRoutes(apiGroup *gin.RouterGroup, c CommentsController) {
	group := apiGroup.Group("/translation-comments/:translationId")
	group.GET("/", c.GetCommentsHandler)
	group.POST("/", c.PostCommentHandler)
	//group.PUT("/:commentID", c.PutCommentHandler)
}

// NewCommentsController is a CommentsController factory function
func NewCommentsController(commentsRepo ports.CommentsRepository, translationRepo ports.TranslationsRepository) CommentsController {
	return CommentsController{
		commentsRepo:     commentsRepo,
		translationdRepo: translationRepo,
	}
}

// GetCommentsHandler is the GET comments handler function
func (ctrl *CommentsController) GetCommentsHandler(c *gin.Context) {
	id := c.Param("translationId")
	comments, err := ctrl.commentsRepo.GetAll(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_type": http.StatusText(http.StatusNotFound),
			"name":       "Resource not found",
			"message":    errors.Wrap(err, fmt.Sprintf("the comments were not found for this translation: %s", id)),
		})
	}

	dtos := lo.Map[domain.Comment, dto.CommentDTO](comments, func(comment domain.Comment, _ int) dto.CommentDTO {
		return dto.ToCommentDTO(comment)
	})
	c.JSON(http.StatusOK, gin.H{
		"comments": dtos,
		"total":    len(dtos),
	})
}

// PostCommentHandler is the POST comments handler function
func (ctrl *CommentsController) PostCommentHandler(c *gin.Context) {
	id := c.Param("translationId")
	var body dto.CreateTranslationCommentRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			"impossible to create a new comment",
		))
		return
	}

	_, err := ctrl.translationdRepo.GetOneForProject(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_type": http.StatusText(http.StatusNotFound),
			"name":       "Resource not found",
			"message":    errors.Wrap(err, fmt.Sprintf("the comment was not found for this translation: %s", id)),
		})
	}

	err = ctrl.commentsRepo.Add(c, id, body.UserId, body.Text)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, buildGinErrorJSON(
			http.StatusBadRequest,
			"Resource not created",
			fmt.Sprintf("impossible to create a new comment for the translation %s", id),
		))
		return
	}

	c.Status(http.StatusCreated)
}
