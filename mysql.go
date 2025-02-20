package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		panic(".envファイルの読み込みに失敗しました")
	}

	var user, password, host, port, name string

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	name = os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || name == "" {
		panic("データベースへの接続に必要な情報が揃っていません")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo", user, password, host, port, name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		println(err)
		panic("データベースの接続に失敗しました")
	}

	println("データベースに接続しました")
}

func GetDB() *gorm.DB {
	return DB
}
