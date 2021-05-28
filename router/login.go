package router

import (
	"net/http"

	"gin-jwt-demo/pkg/code"
	"gin-jwt-demo/pkg/jwt"
	"github.com/gin-gonic/gin"
)

var users = map[string]struct{}{
	"hequan": {},
	"admin":  {},
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": code.ERR_LOGIN_USERNAME,
			"msg":  code.GetMsg(code.ERR_LOGIN_USERNAME),
		})
		return
	}
	if _, ok := users[username]; !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": code.ERR_LOGIN_USERNAME,
			"msg":  code.GetMsg(code.ERR_LOGIN_USERNAME),
		})
		return
	}

	token, err := jwt.GenToken(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": code.FAIL,
			"msg":  code.GetMsg(code.FAIL),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.SUCCESS,
		"msg":  code.GetMsg(code.SUCCESS),
		"data": gin.H{
			"token": token,
		},
	})
}
