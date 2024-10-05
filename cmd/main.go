package main

import (
	"github.com/fs0414/internal/adapter/api/router"
	"github.com/fs0414/internal/infrastructure/database"
)

func main() {
	database.DbInit()
	router := router.GetRouter()
	router.Run(":9090")
}
