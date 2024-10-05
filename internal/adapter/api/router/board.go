package router

import (
	"github.com/gin-gonic/gin"

	repository "github.com/fs0414/internal/adapter/repository/board"
	"github.com/fs0414/internal/usecase"
)

func SetBoardRoutes(rg *gin.RouterGroup) {
	boardUseCase := usecase.BoardUseCaseFactory(repository.BoardRepositoryFactory())

	rg.GET("/boards", boardUseCase.FetchBoards)
	rg.POST("/boards", boardUseCase.CreateBoard)
	rg.DELETE("/boards/:id", boardUseCase.DeleteBoard)
}
