package repository

import (
	"database/sql"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) Create(user domain.User) (domain.User, error) {
	// usersテーブルにデータを追加するクエリ
	var query = `
		INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?);
	`

	// クエリを実行し、userドメインのフィールドを追加
	_, err := repo.db.Exec(query, user.ID(), user.Name(), user.Email(), user.Password())
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo UserRepository) FindByEmail(email string) (*domain.User, error) {
	// usersテーブルからemailフィールドが一致するカラムを取得するクエリ
	var query = `
		SELECT id, name, email, password FROM users WHERE email = ?;
	`

	// emailが一致したカラムを取得
	row := repo.db.QueryRow(query, email)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var dbuser dbUser

	// 取得したカラムから、DB用userモデルの各フィールドに値をスキャン
	err := row.Scan(&dbuser.id, &dbuser.name, &dbuser.email, &dbuser.password, &dbuser.created_at, &dbuser.updated_at)
	if err != nil {
		return nil, err
	}

	// DB用モデルからドメインモデルを生成
	user, err := domain.ReUser(domain.UserID(dbuser.id), dbuser.name, dbuser.email, dbuser.password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
