package main

import (
	"log"
	"time"

	"git.xenonstack.com/check/jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	data := jwt.GinJwtToken()
	log.Println(data)

	log.Println(time.Hour * 24)
	authMiddleware := jwt.MwInitializer()

	r := gin.Default()

	r.Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
