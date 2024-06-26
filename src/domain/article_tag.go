package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type TagID string

type ArticleTag struct {
	id      TagID
	name    string
	usedNum int
}

func NewArticleTag(usedNum int, name string) (*ArticleTag, error) {
	// ulidパッケージでULIDを生成し、string型に変換し、ArticleID型に変換
	id := TagID(ulid.Make().String())

	// 記事タグのバリデーション
	if utf8.RuneCountInString(name) < tagLengthMin || tagLengthMax < utf8.RuneCountInString(name) {
		return nil, errors.New("name is an incorrect value")
	}

	return &ArticleTag{
		id: id,
		name: name,
		usedNum: usedNum,
	}, nil
}

const (
	// 記事タグ名の最小文字数/最大文字数
	tagLengthMin = 1
	tagLengthMax = 20
)
