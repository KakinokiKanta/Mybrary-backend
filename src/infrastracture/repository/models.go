package repository

import (
	"time"
)

type DBUser struct {
	id        string
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}