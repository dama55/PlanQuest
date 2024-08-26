package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "my-go-backend/handlers"
    "my-go-backend/middleware"
    "time"
)


func main() {
    r := gin.Default() // Ginのデフォルト設定でインスタンスを作成

	// ミドルウェアをグローバルに適用
	r.Use(middleware.LoggingMiddleware())

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://frontend:3000", },
        AllowMethods: []string{"GET", "POST"},
        AllowHeaders: []string{"Origin", "Content-Type"},
        ExposeHeaders: []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12* time.Hour,
    }))


    // ルートエンドポイント
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Welcome! The server is running.")
    })

    // "/user"エンドポイントに対するハンドラを設定
    r.GET("/user", handlers.UserHandler)

    // サーバーの起動
    log.Println("Starting server on :8080...")
    r.Run("backend:8080") // デフォルトで ":8080" でサーバーを起動
}