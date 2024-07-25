package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/KakinokiKanta/Mybrary-backend/pkg"
)

type ArticleTagID string

// TODO: 記事タグを機構造にするかも? parentArticleTagID
type ArticleTag struct {
	id      ArticleTagID
	UserID
	tagName string
	usedNum int
}

type ArticleTagRepository interface {
	Create(ArticleTag) (ArticleTag, error)
	FindByName(string) (ArticleTag, error)
	UpdateNum(string) (ArticleTag, error)
}

func NewArticleTag(userID UserID, name string, usedNum int) (*ArticleTag, error) {
	// 記事タグのバリデーション
	if utf8.RuneCountInString(name) < tagLengthMin || tagLengthMax < utf8.RuneCountInString(name) {
		return nil, errors.New("name is an incorrect value")
	}

	return &ArticleTag{
		id: ArticleTagID(pkg.NewULID()),
		UserID: userID,
		tagName: name,
		usedNum: usedNum,
	}, nil
}

func (articleTag ArticleTag) TagName() string {
	return articleTag.tagName
}

func (articleTag ArticleTag) UsedNum() int {
	return articleTag.usedNum
}

const (
	// 記事タグ名の最小文字数/最大文字数
	tagLengthMin = 1
	tagLengthMax = 20
)
