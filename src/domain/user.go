package domain

import (
	"errors"
	"net/mail"
	"unicode/utf8"

	"github.com/KakinokiKanta/Mybrary-backend/pkg"
	"golang.org/x/crypto/bcrypt"
)

type UserID string

type User struct {
	id        UserID
	name      string
	email     string
	password  string
}

type UserRepository interface {
	Create(User) (User, error)
	FindByEmail(string) (User, error)
	// TODO: 現状使わず、未実装のメソッドなのでコメントアウト
	// FindById(UserID) (User, error)
	// Update(User) (User, error)
}

func NewUser(name string, email string, password string) (*User, error) {
	return newUser(
		UserID(pkg.NewULID()),
		name,
		email,
		password,
	)
}

func ReUser(id UserID, name string, email string, password string) (*User, error) {
	return newUser(id, name, email, password)
}

func newUser(id UserID, name string, email string, password string) (*User, error) {
	// IDのバリデーション
	if isValid := pkg.IsValid(string(id)); !isValid {
		return nil, errors.New("id is an incorrect value")
	}

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

	// パスワードをハッシュ化
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		id: id,
		name: name,
		email: addr.Address,
		password: string(hash_pass),
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

func (user User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	return err != bcrypt.ErrMismatchedHashAndPassword
}

const (
	// ユーザ名の最小文字数/最大文字数
	userNameLengthMin = 1
	userNameLengthMax = 20

	// パスワードの最小文字数/最大文字数
	userPassLengthMin = 10
	userPassLengthMax = 64
)
