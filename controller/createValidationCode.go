package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangosteen-go/util"
	"net/http"
)

func CreateValidationCode(ctx *gin.Context) {
	//set header
	ctx.Header("Content-Type", "application/json")
	//get email from request body
	email := ctx.PostForm("email")
	//check email
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email is required",
		})
		return
	}
	// 验证 email 格式
	if !util.CheckEmail(email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email is invalid",
		})
		return
	}
	if err := SendMail(email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "send email failed",
		})
		return
	}
	//发送邮件
	ctx.JSON(http.StatusOK, gin.H{
		"message": "send email success",
	})
}
