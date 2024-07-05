package repository

import (
	"database/sql"
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/infrastracture"
)

// リポジトリテスト全体で共有するsql.DB型
var testDB *sql.DB

func TestMain(m *testing.M)  {
	testDB = infrastracture.SetupTest()

	// パッケージ内のユニットテストをすべて実行
	m.Run()

	// 後処理
	teardown()
}

// 接続したデータベースとのアクセスを閉じる後処理
func teardown() {
	infrastracture.CleanupDB()
	testDB.Close()
}
