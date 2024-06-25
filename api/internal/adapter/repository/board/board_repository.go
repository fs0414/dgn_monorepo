package repository

import (
	"github.com/fs0414/go_hobby/internal/infrastructure/database"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type BoardRepository struct {}

func BoardRepositoryFactory() BoardRepositoryImpl {
    return &BoardRepository{}
}

func (r *BoardRepository) GetBoards() ([]model.Board, error) {
	db := database.GetDb()
	var boards []model.Board
	result := db.Find(&boards)

	if result.Error!= nil {
        return nil, result.Error
    }

	return boards, nil
}

func (repo *BoardRepository) CreateBoard(board model.Board) (model.Board, error) {
	db := database.GetDb()
	result := db.Create(&board)
	if result.Error != nil {
		return model.Board{}, result.Error
	}

	return board, nil
}

func (repo *BoardRepository) DeleteBoard(boardIdInt int) error {
	db := database.GetDb()
    result := db.Delete(&model.Board{}, boardIdInt)
    if result.Error!= nil {
        return result.Error
    }

    return nil
}