package router

import (
	"fmt"
	"net/http"

	"gin-jwt-demo/pkg/code"
	"github.com/gin-gonic/gin"
)

func Something(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": code.FAIL,
			"msg":  code.GetMsg(code.FAIL),
		})
		return
	}
	u, ok := username.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": code.FAIL,
			"msg":  code.GetMsg(code.FAIL),
		})
		return
	}

	fmt.Println("username:", u)
	c.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"msg":  code.GetMsg(code.SUCCESS),
		"data": username,
	})
}
