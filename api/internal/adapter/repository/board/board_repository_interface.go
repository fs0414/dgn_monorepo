package repository

import (
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type BoardRepositoryImpl interface {
	GetBoards() ([]model.Board, error)
	CreateBoard(board model.Board) (model.Board, error)
	DeleteBoard(boardIdInt int) error
}
