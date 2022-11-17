package routes

import (
	"github.com/gin-gonic/gin"

	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
)

// CommentsController is the controller for the projects route
type CommentsController struct {
	commentsRepo ports.CommentsRepository
}

// AddCommentsRoutes adds the routes to the comments endpoint
func AddCommentsRoutes(router *gin.Engine, c CommentsController) {
	group := router.Group("/translation-comments/:translationId")
	group.GET("/", c.GetCommentsHandler)
	group.POST("/", c.PostCommentHandler)
	//group.PUT("/:commentID", c.PutCommentHandler)
}

// NewCommentsController is a CommentsController factory function
func NewCommentsController(repo ports.CommentsRepository) CommentsController {
	return CommentsController{
		commentsRepo: repo,
	}
}

func (ctrl *CommentsController) GetCommentsHandler(c *gin.Context) {
	panic("to implement")
}

func (ctrl *CommentsController) PostCommentHandler(c *gin.Context) {
	panic("to implement")
}
