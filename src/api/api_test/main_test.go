package apitest

import (
	"database/sql"
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/api/router"
	"github.com/KakinokiKanta/Mybrary-backend/infrastracture"
)

// APIテスト全体で共有するsql.DB型
var testDB *sql.DB

func TestMain(m *testing.M)  {
	// 前処理
	testDB = infrastracture.SetupTest()
	router.SetupRouter(testDB)

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
