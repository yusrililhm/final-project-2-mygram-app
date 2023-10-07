package comment_pg

import (
	"database/sql"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/comment_repository"
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
			id, message, photo_id, user_id, created_at
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

	var comment dto.NewCommentResponse
	err = tx.QueryRow(
		addCommentQuery,
		commentPayload.UserId,
		commentPayload.PhotoId,
		commentPayload.Message,
	).Scan(
		&comment.Id,
		&comment.Message,
		&comment.PhotoId,
		&comment.UserId,
		&comment.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &comment, nil
}
