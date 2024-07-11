package domain

import (
	"errors"
	"net/mail"
	"time"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type UserID string

type User struct {
	id        UserID
	name      string
	email     string
	password  string
	createdAt time.Time
}

type UserRepository interface {
	Create(User) (User, error)
	// TODO: 現状使わず、未実装のメソッドなのでコメントアウト
	// FindById(UserID) (User, error)
	// Update(User) (User, error)
}

func NewUser(name string, email string, password string) (*User, error) {
	// ULIDを生成し、string型に変換し、UserID型に変換
	id := UserID(ulid.Make().String())

	// 現在時刻を、記事ドメイン生成時刻とする
	createdAt := time.Now()

	// ユーザ名のバリデーション
	if utf8.RuneCountInString(name) < userNameLengthMin || userNameLengthMax < utf8.RuneCountInString(name) {
		return nil, errors.New("name is an incorrect value")
	}

	// メールアドレスのバリデーション
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	// パスワードのバリデーション
	if utf8.RuneCountInString(password) < userPassLengthMin || userPassLengthMax < utf8.RuneCountInString(password) {
		return nil, errors.New("password is an incorrect value")
	}

	return &User{
		id: id,
		name: name,
		email: addr.Address,
		password: password,
		createdAt: createdAt,
	}, nil
}

func (user User) ID() UserID {
	return user.id
}

func (user User) Name() string {
	return user.name
}

func (user User) Email() string {
	return user.email
}

func (user User) Password() string {
	return user.password
}

func (user User) CreatedAt() time.Time {
	return user.createdAt
}

const (
	// ユーザ名の最小文字数/最大文字数
	userNameLengthMin = 1
	userNameLengthMax = 20

	// パスワードの最小文字数/最大文字数
	userPassLengthMin = 10
	userPassLengthMax = 64
)
