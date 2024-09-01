package repository

import (
	"database/sql"
	"strconv"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type ArticleTagRepository struct {
	db *sql.DB
}

func NewArticleTagRepository(db *sql.DB) ArticleTagRepository {
	return ArticleTagRepository{
		db: db,
	}
}

// tagsテーブルにレコードを追加するメソッド
func (repo ArticleTagRepository) Create(articleTag domain.ArticleTag) (domain.ArticleTag, error) {
	// tagsテーブルにデータを追加するクエリ
	var query = `
		INSERT INTO tags (id, user_id, name, used_num) VALUES (?, ?, ?, ?);
	`

	// クエリを実行し、userドメインのフィールドを追加
	_, err := repo.db.Exec(query, articleTag.ID(), articleTag.UserID, articleTag.TagName(), articleTag.UsedNum())
	if err != nil {
		return domain.ArticleTag{}, err
	}

	return articleTag, nil
}

// tagsテーブルからnameフィールドが一致するレコードを取得するメソッド
func (repo ArticleTagRepository) FindByName(name string) (domain.ArticleTag, error) {
	// tagsテーブルからnameフィールドが一致するレコードを取得するクエリ
	var query = `
		SELECT id, user_id, name, used_num FROM tags WHERE name = ?;
	`

	// nameが一致したレコードを取得
	row := repo.db.QueryRow(query, name)
	if err := row.Err(); err != nil {
		return domain.ArticleTag{}, err
	}

	var dbTag dbTag

	// 取得したレコードから、DB用tagモデルの各フィールドに値をスキャン
	err := row.Scan(&dbTag.id, &dbTag.user_id, &dbTag.name, &dbTag.used_num)
	if err != nil {
		return domain.ArticleTag{}, err
	}

	// DB用モデルからドメインモデルを生成
	num, err := strconv.Atoi(dbTag.used_num)
	if err != nil {
		return domain.ArticleTag{}, err
	}
	articleTag, err := domain.ReArticleTag(domain.ArticleTagID(dbTag.id), domain.UserID(dbTag.user_id), dbTag.name, num)
	if err != nil {
		return domain.ArticleTag{}, err
	}

	return *articleTag, nil
}

// tagsテーブルのuser_idフィールドが一致するすべてのレコードを取得するメソッド
func (repo ArticleTagRepository) FindByUserID(id domain.UserID) ([]domain.ArticleTag, error) {
	return nil, nil
}

// tagsテーブルのidフィールドが一致するレコードのused_numを1増やすメソッド
func (repo ArticleTagRepository) UpdateNum(id domain.ArticleTagID) error {
	// tagsテーブルのidフィールドが一致するレコードのused_numを1増やすクエリ
	var query = `
		UPDATE tags SET used_num = used_num + 1 WHERE id = ?;
	`

	// クエリを実行
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
