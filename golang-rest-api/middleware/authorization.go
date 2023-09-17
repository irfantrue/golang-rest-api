package middleware

import (
	"golang-rest-api/helper"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if "RAHASIA" != c.GetHeader("Authorization") {
			helper.MiddlewareResponse(c, 401)
		}
		c.Next()
	}
}
