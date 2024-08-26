package middleware

import (
	"log"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		// リクエストのログを出力
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		// 次のミドルウェアやハンドラを実行
		c.Next()
		// レスポンス後のログを出力（必要なら）
		log.Printf("Response status: %d", c.Writer.Status())
	}
}

