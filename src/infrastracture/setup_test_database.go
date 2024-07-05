package infrastracture

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	testDBHost string
	testDBUser string
	testDBPassword string
	testDBName string
	testDBPort string
)

func SetupTestDB() *sql.DB {
	// 環境変数からデータベース接続情報を取得
	loadTestEnv()

	// データベース接続文字列の構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		testDBUser, testDBPassword, testDBHost, testDBPort, testDBName)

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

func loadTestEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("fail to load .env file")
	}
	testDBHost = os.Getenv("TEST_DB_HOST")
	testDBUser = os.Getenv("TEST_MYSQL_USER")
	testDBPassword = os.Getenv("TEST_MYSQL_PASSWORD")
	testDBName = os.Getenv("TEST_MYSQL_DATABASE")
	testDBPort = os.Getenv("TEST_DB_PORT")
}
