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
		INSERT INTO users (id, name, email, password, created_at) VALUES (?, ?, ?, ?, ?);
	`
	// sql.DB型のメソッドExecを用いて、クエリを実行
	_, err := repo.db.Exec(query, user.ID(), user.Name(), user.Email(), user.Password(), user.CreatedAt())
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo UserRepository) FindByEmail(user domain.User) (*domain.User, error) {
	// usersテーブルからemailフィールドが一致するカラムを取得するクエリ
	var query = `
		SELECT * FROM users WHERE email = ?;
	`
}
