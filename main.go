package main

import (
	"practice-api/infrastructure/router"
	"practice-api/internal/user/handler"
	"practice-api/internal/user/repository"
	"practice-api/internal/user/service"

	"gorm.io/gorm"
)

func main() {
	// DBの初期化
	var database *gorm.DB
	Connect()
	database = GetDB()
	// リポジトリ、バリデータ、サービス、ハンドラの初期化
	userRepository := userRepository.NewUserRepository(database)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	// ルーターの設定
	r := router.SetUpUserRoutes(userHandler)

	if err := r.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
