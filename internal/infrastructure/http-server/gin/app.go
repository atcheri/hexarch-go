package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/atcheri/hexarch-go/docs"
	dto "github.com/atcheri/hexarch-go/internal/core/dtos"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
)

//// GinApp will implement the Server interface
//type GinApp struct{}

func NewGinApp(wordsController WordsController, _ ports.SentencesRepository) *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"date": time.Now().String()})
	})
	wordsRoutes(app, wordsController)
	app.StaticFS("/api/docs", http.FS(docs.Swagger))

	return app
}

func wordsRoutes(router *gin.Engine, wh WordsController) {
	wordsGroup := router.Group("/words")
	wordsGroup.GET("/", wh.GetAllHandler)
	wordsGroup.POST("/", wh.CreateWordHandler)
	wordsGroup.PUT("/", wh.UpdateWordHandler)
	wordsGroup.DELETE("/", wh.DeleteWordHandler)
}

type WordsController struct {
	wordsRepo ports.WordsRepository
}

func NewWordsController(wordsRepo ports.WordsRepository) WordsController {
	return WordsController{
		wordsRepo: wordsRepo,
	}
}

func (wh WordsController) GetAllHandler(c *gin.Context) {
	offset := 0
	limit := 5
	c.JSON(http.StatusOK, gin.H{"words": wh.wordsRepo.GetAll(offset, limit)})
}

func (wh WordsController) CreateWordHandler(c *gin.Context) {
	body := dto.WordsPostRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{
			ErrorType: "BadPostRequestBody-type",
			Message:   "Invalid request body",
			Name:      "Bad POST Request body",
		})
		return
	}
	wh.wordsRepo.SetWord(body.Key, body.Content)
	c.JSON(http.StatusCreated, gin.H{"id": body.Key})
}

func (wh WordsController) UpdateWordHandler(c *gin.Context) {
	body := dto.WordsPostRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{
			ErrorType: "BadPutRequestBody-type",
			Message:   "Invalid request body",
			Name:      "Bad PUT Request body",
		})
		return
	}
	wh.wordsRepo.SetWord(body.Key, body.Content)
	c.JSON(http.StatusOK, gin.H{"id": body.Key})
}

func (wh WordsController) DeleteWordHandler(c *gin.Context) {
	body := dto.WordsDeleteRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{
			ErrorType: "BadDeleteRequestBody-type",
			Message:   "Invalid request body",
			Name:      "Bad DELETE Request body",
		})
		return
	}
	if err := wh.wordsRepo.RemoveWord(body.Key); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{
			ErrorType: "BadDeleteRequestBody-type",
			Message:   "Could not delete the word",
			Name:      "Bad DELETE Request body",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": "word deleted"})
}
