package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func guidMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fetch data from headers

		authorization := ctx.GetHeader("Authorization")

		// get params from URL

		name := ctx.Param("name")

		fmt.Println(authorization, name)

		if name == "" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid Request!",
			})

			ctx.Abort()
			return

		}

		ctx.Next()

	}
}

func main() {

	router := gin.Default()

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

	})

	router.Run(":8081")

}
