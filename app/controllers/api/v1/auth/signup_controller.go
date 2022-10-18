package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/vinoMamba/mangosteen-go/app/controllers/api/v1"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	type EmailExistRequest struct {
		Email string `json:"email"`
	}
	request := EmailExistRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data": request.Email,
		})
	}
}
