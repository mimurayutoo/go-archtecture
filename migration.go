package main

// import (
// 	"fmt"
// 	"log"
// 	"practice-api/infrastructure/db"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("Error loading .env file")
// 	}

// 	db.Connect()
// 	database := db.GetDB()
// 	if database == nil {
// 		log.Fatalf("Failed to connect to the database")
// 	}

// 	fmt.Println("DB マイグレーション開始")
// 	db.Migrate(database)
// 	fmt.Println("DB マイグレーション完了")
// }
