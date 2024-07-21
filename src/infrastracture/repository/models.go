package repository

import (
	"database/sql"
)

/*
 * ドメイン層のモデルに対して、このファイルでは、
 * データベースのテーブルを構造体としたモデルを定義する
**/
type dbUser struct {
	id         string
	name       string
	email      string
	password   string
	created_at sql.NullTime
	updated_at sql.NullTime
}