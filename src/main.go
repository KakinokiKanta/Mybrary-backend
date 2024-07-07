package main

import (
	"github.com/KakinokiKanta/Mybrary-backend/api/router"
	"github.com/KakinokiKanta/Mybrary-backend/infrastracture"
)

func main() {
	db := infrastracture.SetupDB()
	defer db.Close()
	router.SetupRouter(db)
}
