package main

import (
  "github.com/fs0414/go_hobby/internal/infrastructure/database"
  "github.com/fs0414/go_hobby/internal/adapter/api/router"
)

func main() {
  database.DbInit()
  router := router.GetRouter()
  router.Run(":8080")
}