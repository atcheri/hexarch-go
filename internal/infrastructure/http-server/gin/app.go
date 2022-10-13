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

func NewGinApp(wordsRepo ports.WordsRepository, _ ports.SentencesRepository) *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"date": time.Now().String()})
	})

	wordsGroup := app.Group("/words")
	wordsGroup.GET("", func(c *gin.Context) {
		offset := 0
		limit := 5
		c.JSON(http.StatusOK, gin.H{"words": wordsRepo.GetAll(offset, limit)})
	})
	wordsGroup.POST("", func(c *gin.Context) {
		body := dto.WordsPostRequestBody{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.Error{
				ErrorType: "BadRequestBody-type",
				Message:   "Invalid request body",
				Name:      "Bad Request body",
			})
			return
		}
		wordsRepo.SetWord(body.Key, body.Content)
		c.JSON(http.StatusOK, gin.H{"id": body.Key})
	})

	app.StaticFS("/api/docs", http.FS(docs.Swagger))

	return app
}
