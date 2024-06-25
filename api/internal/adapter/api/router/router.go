package router
// "github.com/fs0414/go_hobby/internal/adapter/api/router"

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/", func (c *gin.Context)  {
    c.String(200, "hello go hobby")
  })

  apiRouter := r.Group("/api")
  {
    SetUserRoutes(apiRouter)
    SetBoardRoutes(apiRouter)
  }

  return r
}