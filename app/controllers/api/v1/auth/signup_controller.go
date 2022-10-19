package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/vinoMamba/mangosteen-go/app/controllers/api/v1"
	"github.com/vinoMamba/mangosteen-go/app/requests"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.EmailExistRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}
	requests.ValidateEmailExistRequest(request)
	c.JSON(200, gin.H{
		"data": request.Email,
	})
}
