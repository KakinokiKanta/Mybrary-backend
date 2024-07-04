package main

import (
	"github.com/KakinokiKanta/Mybrary-backend/api/router"
	database "github.com/KakinokiKanta/Mybrary-backend/infrastracture"
)

func main() {
	db := database.SetupDB()
	router.SetupRouter(db)
}
