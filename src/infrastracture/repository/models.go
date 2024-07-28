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

type dbTag struct {
	id         string
	user_id    string
	tag_name   string
	used_num   string
	created_at sql.NullTime
	updated_at sql.NullTime
}
