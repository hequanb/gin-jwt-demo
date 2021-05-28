package main

import (
	"gin-jwt-demo/middleware/jwt"
	"gin-jwt-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", router.Login)

	authorized := r.Group("/auth")
	authorized.Use(jwt.JWTAuthMiddleware())
	{
		authorized.GET("/something", router.Something)
	}

	r.Run(":8080")
}
