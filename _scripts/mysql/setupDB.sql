# ユーザデータを格納するためのテーブル
CREATE TABLE users (
  id VARCHAR(128) NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
)
