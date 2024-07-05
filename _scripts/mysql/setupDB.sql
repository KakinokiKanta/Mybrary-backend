# ユーザデータを格納するためのテーブル
CREATE TABLE users (
  id VARCHAR(128) NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  created_at datetime,
  update_at datetime
)
