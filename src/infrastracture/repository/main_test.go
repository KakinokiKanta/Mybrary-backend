package repository

import (
	"database/sql"
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	"github.com/KakinokiKanta/Mybrary-backend/infrastracture"
)

// リポジトリテスト全体で共有
var testDB *sql.DB
var succeedDomain *domain.ArticleTag
const succeedTagName = "New Article Tag"

func TestMain(m *testing.M)  {
	// 前処理
	setup()

	// パッケージ内のユニットテストをすべて実行
	m.Run()

	// 後処理
	teardown()
}

// 前処理
func setup() {
	testDB = infrastracture.SetupTest()
	succeedDomain, _ = domain.NewArticleTag("abcdefg1234AABBCCDD", succeedTagName)
}

// 接続したデータベースとのアクセスを閉じる後処理
func teardown() {
	infrastracture.CleanupDB()
	testDB.Close()
}
