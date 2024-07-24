package domain

import (
	"errors"
	"unicode/utf8"
)

type ArticleTagID string

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

func NewArticleTag(name string, usedNum int) (*ArticleTag, error) {
	// 記事タグのバリデーション
	if utf8.RuneCountInString(name) < tagLengthMin || tagLengthMax < utf8.RuneCountInString(name) {
		return nil, errors.New("name is an incorrect value")
	}

	return &ArticleTag{
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
