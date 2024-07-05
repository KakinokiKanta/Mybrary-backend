package infrastracture

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

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

func SetupTest() *sql.DB {
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

	// テスト用データの登録
	if err = CleanupDB(); err != nil {
		fmt.Println("Cleanup database")
		return nil
	}
	if err = setupTestData(); err != nil {
		fmt.Println("Setup test data")
		return nil
	}

	return db
}

func loadTestEnv() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("fail to load .env file")
	}
	testDBHost = os.Getenv("TEST_DB_HOST")
	testDBUser = os.Getenv("TEST_MYSQL_USER")
	testDBPassword = os.Getenv("TEST_MYSQL_PASSWORD")
	testDBName = os.Getenv("TEST_MYSQL_DATABASE")
	testDBPort = os.Getenv("TEST_DB_PORT")
}

// DB内のテスト用データを消す後処理
func CleanupDB() error {
	// os/execパッケージのexec.Command関数を用いて、実行したいコマンドの情報を持つexec.Cmd型の変数を用意
	passStr := fmt.Sprintf("--password=%s", testDBPassword)
	cmd := exec.Command("mysql", "-h", testDBHost, "-u", testDBUser, testDBName, passStr, "-e", "source ../../_scripts/mysql/cleanupDB.sql")
	// exec.Cmd型のRunメソッドを読んで、コマンドを実行
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// データベースにデータを入れる前処理
func setupTestData() error {
	// os/execパッケージのexec.Command関数を用いて、実行したいコマンドの情報を持つexec.Cmd型の変数を用意
	passStr := fmt.Sprintf("--password=%s", testDBPassword)
	cmd := exec.Command("mysql", "-h", testDBHost, "-u", testDBUser, testDBName, passStr, "-e", "source ../../_scripts/mysql/setupDB.sql")
	// exec.Cmd型のRunメソッドを読んで、コマンドを実行
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	cmd = exec.Command("mysql", "-h", testDBHost, "-u", testDBUser, testDBName, passStr, "-e", "source ../../_scripts/mysql/insertUser.sql")
	// exec.Cmd型のRunメソッドを読んで、コマンドを実行
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
