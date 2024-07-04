# ユーザデータを格納するためのテーブル
CREATE TABLE user (
  id VARCHAR(128) NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  createdAt datetime
)
