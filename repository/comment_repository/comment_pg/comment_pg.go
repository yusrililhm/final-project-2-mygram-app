package comment_pg

import (
	"database/sql"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/comment_repository"
	"time"
)

type commentRepositoryImpl struct {
	db *sql.DB
}

const (
	addCommentQuery = `
		INSERT INTO
			comments
				(
					user_id,
					photo_id,
					message
				)
		VALUES
				(
					$1, 
					$2, 
					$3
				)
		RETURNING
			id
	`
)

func NewCommentRepository(db *sql.DB) comment_repository.CommentRepository {
	return &commentRepositoryImpl{
		db: db,
	}
}

// AddComment implements comment_repository.CommentRepository.
func (commentRepo *commentRepositoryImpl) AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error) {
	tx, err := commentRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	var id int
	err = tx.QueryRow(addCommentQuery, commentPayload.UserId, commentPayload.PhotoId, commentPayload.Message).Scan(&id)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &dto.NewCommentResponse{
		Id:        id,
		UserId:    commentPayload.UserId,
		PhotoId:   commentPayload.PhotoId,
		Message:   commentPayload.Message,
		CreatedAt: time.Now(),
	}, nil
}
