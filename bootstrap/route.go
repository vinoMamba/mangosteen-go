package bootstrap

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/mangosteen-go/routes"
)

func SetupRouter(r *gin.Engine) {
	registerMiddleware(r)
	routes.RegisterApiRoutes(r)
	registerNotFoundRoute(r)
}

func registerMiddleware(r *gin.Engine) {
	// Register middleware
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}
func registerNotFoundRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			c.String(404, "404 page not found")
		} else {
			c.JSON(404, gin.H{
				"code":    404,
				"message": "404 page not found",
			})
		}
	})
}
