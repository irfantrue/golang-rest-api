package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // Lanjutkan penanganan permintaan
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Mencetak log dengan warna menggunakan fmt
		ip := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		userAgent := c.Request.UserAgent()

		// Mencetak log dengan warna
		fmt.Printf("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("IP: %s, ", ip)
		fmt.Printf("Latency: %v, ", latency)
		fmt.Printf("Method: %s, ", method)
		fmt.Printf("Path: %s, ", path)
		fmt.Printf("Status: %d, ", status)
		fmt.Printf("User-Agent: %s\n", userAgent)
	}
}
