package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	dbHost string
	dbUser string
	dbPassword string
	dbName string
	dbPort string
)

func SetupDB() *sql.DB {
	// 環境変数からデータベース接続情報を取得
	loadEnv()

	// データベース接続文字列の構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// データベースに接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// データベース接続の確認
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("Succeeded to connect database")

	return db
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("fail to load .env file")
	}
	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbName = os.Getenv("MYSQL_DATABASE")
	dbPort = os.Getenv("DB_PORT")
}

