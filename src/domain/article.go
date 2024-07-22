package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/KakinokiKanta/Mybrary-backend/pkg"
)

type ArticleID string

type Article struct {
	id          ArticleID
	userID      UserID
	url         string
	title       string
	description string
	createdAt   time.Time
	tags        []ArticleTag
}

type ArticleRepository interface {
	Create(Article) (Article, error)
	FindByArticleId(ArticleID) (Article, error)
	FindByUserID(UserID) ([]Article, error)
	Update(Article) (Article, error)
	Delete(Article) error
}

func NewArticle(
	userID UserID,
	url string,
	title string,
	description string,
	tags []ArticleTag,
) (*Article, error) {
	// ulidパッケージでULIDを生成し、string型に変換し、ArticleID型に変換
	id := ArticleID(pkg.NewULID())

	// timeパッケージで現在時刻を、記事ドメイン生成時刻とする
	createdAt := time.Now()

	// URLのバリデーション
	if utf8.RuneCountInString(url) < urlLengthMin {
		return nil, errors.New("url is an incorrect value")
	}
	// 記事タイトルのバリデーション
	if utf8.RuneCountInString(title) < titleLengthMin || titleLengthMax < utf8.RuneCountInString(title) {
		return nil, errors.New("title is an incorrect value")
	}
	// 記事詳細のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || descriptionLengthMax < utf8.RuneCountInString(description) {
		return nil, errors.New("description is an incorrect value")
	}
	// 記事に付加するタグのバリデーション
	if len(tags) < tagsLengthMin || tagsLengthMax < len(tags) {
		return nil, errors.New("url is an incorrect value")
	}

	return &Article{
		id: id,
		userID: userID,
		url: url,
		title: title,
		description: description,
		createdAt: createdAt,
		tags: tags,
	}, nil
}

func (article Article) ID() ArticleID {
	return article.id
}

func (article Article) UserID() UserID {
	return article.userID
}

func (article Article) URL() string {
	return article.url
}

func (article Article) Title() string {
	return article.title
}

func (article Article) Description() string {
	return article.description
}

func (article Article) CreatedAt() time.Time {
	return article.createdAt
}

func (article Article) Tags() []ArticleTag {
	return article.tags
}

const (
	// URLの最小値
	urlLengthMin = 1

	// 記事タイトルの最小値/最大値
	titleLengthMin = 1
	titleLengthMax = 100

	// 記事詳細の最小値/最大値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000

	// タグの最小個数/最大個数
	tagsLengthMin = 0
	tagsLengthMax = 5
)
