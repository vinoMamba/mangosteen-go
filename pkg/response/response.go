package response

import "github.com/gin-gonic/gin"

func JSON(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}

// put post delete
func Success(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func Data(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
