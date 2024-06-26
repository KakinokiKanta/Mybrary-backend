package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type ArticleTag struct {
	tagID TagID
	name string
	usedNum int
}

type TagID string

func NewArticleTag(usedNum int, name string) (*ArticleTag, error) {
	// ulidパッケージでULIDを生成し、string型に変換し、ArticleID型に変換
	id := TagID(ulid.Make().String())

	// 記事タグのバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || nameLengthMax < utf8.RuneCountInString(name) {
		return nil, errors.New("name is an incorrect value")
	}

	return &ArticleTag{
		tagID: id,
		name: name,
		usedNum: usedNum,
	}, nil
}

const (
	// 記事タグ名の最小文字数/最大文字数
	nameLengthMin = 1
	nameLengthMax = 20
)
