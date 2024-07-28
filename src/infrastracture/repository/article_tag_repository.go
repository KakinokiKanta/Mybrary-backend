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

func (repo ArticleTagRepository) Create(articleTag domain.ArticleTag) (domain.ArticleTag, error) {
	// tagsテーブルにデータを追加するクエリ
	var query = `
		INSERT INTO tags (id, user_id, tag_name, used_num) VALUES (?, ?, ?, ?);
	`

	// クエリを実行し、userドメインのフィールドを追加
	_, err := repo.db.Exec(query, articleTag.ID(), articleTag.UserID, articleTag.TagName(), articleTag.UsedNum())
	if err != nil {
		return domain.ArticleTag{}, err
	}

	return articleTag, nil
}

func (repo ArticleTagRepository) FindByName(name string) (domain.ArticleTag, error) {
	// tagsテーブルからemailフィールドが一致するカラムを取得するクエリ
	var query = `
		SELECT id, user_id, tag_name, used_num FROM tags WHERE tag_name = ?;
	`

	// nameが一致したカラムを取得
	row := repo.db.QueryRow(query, name)
	if err := row.Err(); err != nil {
		return domain.ArticleTag{}, err
	}

	var dbTag dbTag

	// 取得したカラムから、DB用tagモデルの各フィールドに値をスキャン
	err := row.Scan(&dbTag.id, &dbTag.user_id, &dbTag.tag_name, &dbTag.used_num)
	if err != nil {
		return domain.ArticleTag{}, err
	}

	// DB用モデルからドメインモデルを生成
	num, err := strconv.Atoi(dbTag.used_num)
	if err != nil {
		return domain.ArticleTag{}, err
	}
	articleTag, err := domain.ReArticleTag(domain.ArticleTagID(dbTag.id), domain.UserID(dbTag.user_id), dbTag.tag_name, num)
	if err != nil {
		return domain.ArticleTag{}, err
	}

	return *articleTag, nil
}
