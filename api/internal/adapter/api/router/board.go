package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fs0414/go_hobby/internal/usecase"
	"github.com/fs0414/go_hobby/internal/adapter/repository/board"
)

func SetBoardRoutes(rg *gin.RouterGroup) {
    boardUseCase := usecase.BoardUseCaseFactory(repository.BoardRepositoryFactory())

	rg.GET("/boards", boardUseCase.FetchBoards)
	rg.POST("/boards", boardUseCase.CreateBoard)
	rg.DELETE("/boards/:id", boardUseCase.DeleteBoard)
}