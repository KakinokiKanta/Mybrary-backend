# ユーザデータを格納するためのテーブル
# TODO: created_atとupdate_atはデータベースの設定によって自動で埋めるようにしたい
CREATE TABLE users (
  id VARCHAR(128) NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL,
  password VARCHAR(128) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
