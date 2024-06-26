package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type ArticleID string

type Article struct {
	id          ArticleID
	userID      string
	url         string
	title       string
	description string
	tags        []ArticleTag
}

type ArticleRepository interface {
	Create(Article) error
	FindById(ArticleID) (Article, error)
	FindAll() ([]Article, error)
	Update(Article) error
	Delete(Article) error
}

func NewArticle(
	userID string,
	url string,
	title string,
	description string,
	tags []ArticleTag,
) (*Article, error) {
	// ulidパッケージでULIDを生成し、string型に変換し、ArticleID型に変換
	id := ArticleID(ulid.Make().String())

	// ユーザIDのバリデーション
	if _, err := ulid.Parse(userID); err != nil {
		return nil, err
	}
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
		tags: tags,
	}, nil
}

type ArticleTag struct {
	tagID string
}

func NewArticleTag(tagID string) (*ArticleTag, error) {
	// 記事タグIDのバリデーション
	if _, err := ulid.Parse(tagID); err != nil {
		return nil, err
	}

	return &ArticleTag{
		tagID: tagID,
	}, nil
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
