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
		SELECT * FROM users WHERE email = ?;
	`

	// emailが一致したカラムを取得
	row := repo.db.QueryRow(query, email)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var user DBUser
	var createdAt sql.NullTime
	var updatedAt sql.NullTime

	// 取得したカラムから、DB用userモデルの各フィールドに値をスキャン
	err := row.Scan(&user.id, &user.name, &user.email, &user.password, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	// created_atフィールドがnullでないならば
	if createdAt.Valid {
		user.createdAt = createdAt.Time
	}
	// updated_atフィールドがnullでないならば
	if updatedAt.Valid {
		user.updatedAt = updatedAt.Time
	}

	return 
}
