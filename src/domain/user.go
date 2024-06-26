package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type UserID string

type User struct {
	id        UserID
	name      string
	createdAt time.Time
}

// TODO: リポジトリ
type UserRepository interface {
	
}

func NewUser(name string) (*User, error) {
	// ulidパッケージでULIDを生成し、string型に変換し、UserID型に変換
	id := UserID(ulid.Make().String())

	// timeパッケージで現在時刻を、記事ドメイン生成時刻とする
	createdAt := time.Now()

	// ユーザ名のバリデーション
	if utf8.RuneCountInString(name) < userNameLengthMin || userNameLengthMax < utf8.RuneCountInString(name) {
		return nil, errors.New("name is an incorrect value")
	}

	return &User{
		id: id,
		name: name,
		createdAt: createdAt,
	}, nil
}

const (
	// ユーザ名の最小文字数/最大文字数
	userNameLengthMin = 1
	userNameLengthMax = 20
)
