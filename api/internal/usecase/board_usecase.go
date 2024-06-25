package usecase

import (
	// "fmt"
	// "reflect"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fs0414/go_hobby/internal/adapter/repository/board"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type BoardUseCase struct {
	repo repository.BoardRepositoryImpl
}

type BoardCreateRequest struct {
	Content string `json:"content" binding:"required,min=1,max=255"`
	UserID uint `json:"user_id"`
}

func BoardUseCaseFactory(repo repository.BoardRepositoryImpl) *BoardUseCase {
	return &BoardUseCase{repo: repo}
}

func (uc *BoardUseCase) FetchBoards(c *gin.Context) {
	boards, err := uc.repo.GetBoards()
    if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, boards)
}

func (uc *BoardUseCase) CreateBoard(c *gin.Context) {
	var req BoardCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	board := model.Board{
        Content: req.Content,
        UserID: req.UserID,
    }

	newBoard, err := uc.repo.CreateBoard(board)

	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

	c.JSON(201, newBoard)
}

func (uc *BoardUseCase)DeleteBoard(c *gin.Context) {
	boardId := c.Param("id")
	boardIdUint, err := strconv.Atoi(boardId)
	
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	
	if err := uc.repo.DeleteBoard(boardIdUint); err!= nil {
		c.JSON(500, gin.H{"error": err.Error()})
        return
	}

	c.Status(204)
}