package repository

import (
	"database/sql"

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
	// articleTagテーブルにデータを追加するクエリ
	var query = `
		INSERT INTO article_tags (id, user_id, tag_name, used_num) VALUES (?, ?, ?, ?);
	`

	// クエリを実行し、userドメインのフィールドを追加
	_, err := repo.db.Exec(query, articleTag.ID(), articleTag.UserID, articleTag.TagName(), articleTag.UsedNum())
	if err != nil {
		return domain.ArticleTag{}, err
	}

	return articleTag, nil
}
