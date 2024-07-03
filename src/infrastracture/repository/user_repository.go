package repository

import "github.com/KakinokiKanta/Mybrary-backend/domain"

// TODO: ここにデータベース処理を実装して
// データベースを操作するExec処理等は
// databaseに移行する
type UserRepository struct {}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (repo UserRepository) Create(user domain.User) (domain.User, error) {
	// クエリの定義
	// var query = `
	// 	INSERT INTO users (id, name, created_at) values (?, ?, now());
	// `
	// sql.DB型のメソッドExecを用いて、クエリを実行
	// result, err := 

	return domain.User{}, nil // いったんね
}