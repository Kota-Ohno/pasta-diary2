package main

import (
	"log"
	"pasta-diary2-backend/config"
	"pasta-diary2-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // MongoDB接続の初期化
    config.ConnectDatabase()

    // ルーティングの設定
    routes.SetupRoutes(router)

    // サーバーの起動
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
