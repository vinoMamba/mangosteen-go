package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/mangosteen-go/app/controllers/api/v1/auth"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}
}
