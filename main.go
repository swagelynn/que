package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memoryStore "github.com/ulule/limiter/v3/drivers/store/memory"
)

type Question struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func makeLimiterMiddleware(rate limiter.Rate) gin.HandlerFunc {
	store := memoryStore.NewStore()
	instance := limiter.New(store, rate)
	return ginlimiter.NewMiddleware(instance)
}

func main() {
	ensureFiles()

	g := gin.Default()

	g.POST("/question", makeLimiterMiddleware(limiter.Rate{Period: time.Minute, Limit: 1}), func(ctx *gin.Context) {
		var body Question

		ctx.ShouldBindJSON(&body)

		ctx.JSON(writeQuestion(body), gin.H{})
	})

	g.GET("/questions", makeLimiterMiddleware(limiter.Rate{Period: time.Minute, Limit: 50}), func(ctx *gin.Context) {
		ctx.String(200, string(getDisplayData()))
	})

	g.Run(":5438")
}
