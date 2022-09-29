package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//ValidationCodeRouter(r)
	v1 := r.Group("/api/v1")
	validationCodes := v1.Group("/validation_codes")
	{
		validationCodes.POST("/", func(ctx *gin.Context) {})
	}
	session := v1.Group("/session")
	{
		session.POST("/", func(ctx *gin.Context) {})
		session.DELETE("/", func(ctx *gin.Context) {})
	}
	items := v1.Group("/items")
	{
		items.GET("/", func(ctx *gin.Context) {})
		items.POST("/", func(ctx *gin.Context) {})
		items.GET("/:id", func(ctx *gin.Context) {})
		items.PUT("/:id", func(ctx *gin.Context) {})
		items.DELETE("/:id", func(ctx *gin.Context) {})
	}
	tags := v1.Group("/tags")
	{
		tags.GET("/", func(ctx *gin.Context) {})
		tags.POST("/", func(ctx *gin.Context) {})
		tags.GET("/:id", func(ctx *gin.Context) {})
		tags.PUT("/:id", func(ctx *gin.Context) {})
		tags.DELETE("/:id", func(ctx *gin.Context) {})
	}
	return r
}
