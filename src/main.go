package main

import (
	"github.com/KakinokiKanta/Mybrary-backend/api/router"
	database "github.com/KakinokiKanta/Mybrary-backend/infrastracture"
)

func main() {
	db := database.SetupDB()
	r := router.SetupRouter(db)

	r.Run() // dockerでポート8080を指定しているため、ここでは指定しない
}
