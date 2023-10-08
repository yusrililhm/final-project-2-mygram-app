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

	getCommentQuery = `
		SELECT 
			c.id,
			c.user_id,
			c.photo_id,
			c.message,
			c.created_at,
			c.updated_at,
			u.id,
			u.username,
			u.email,
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id
		FROM 
			comments AS c
		LEFT JOIN
			users AS u
		ON
			c.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			c.photo_id = p.id
		ORDER BY 
			c.id
		ASC
	`

	getCommentByIdQuery = `
		SELECT 
			c.id,
			c.user_id,
			c.photo_id,
			c.message,
			c.created_at,
			c.updated_at,
			u.id,
			u.username,
			u.email,
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id
		FROM 
			comments AS c
		LEFT JOIN
			users AS u
		ON
			c.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			c.photo_id = p.id
		WHERE c.id = $1
	`

	deleteCommentQuery = `
		DELETE FROM
			comments
		WHERE
			id = $1
	`

	updateCommentQuery = `
		UPDATE 
			comments
		SET
			message = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, user_id, photo_id, message, updated_at
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

// GetComments implements comment_repository.CommentRepository.
func (commentRepo *commentRepositoryImpl) GetComments() ([]comment_repository.CommentUserPhotoMapped, errs.Error) {

	var commentsUserPhoto []comment_repository.CommentUserPhoto
	rows, err := commentRepo.db.Query(getCommentQuery)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	for rows.Next() {
		commentUserPhoto := comment_repository.CommentUserPhoto{}
		err = rows.Scan(
			&commentUserPhoto.Comment.Id,
			&commentUserPhoto.Comment.UserId,
			&commentUserPhoto.Comment.PhotoId,
			&commentUserPhoto.Comment.Message,
			&commentUserPhoto.Comment.CreatedAt,
			&commentUserPhoto.Comment.UpdatedAt,
			&commentUserPhoto.User.Id,
			&commentUserPhoto.User.Username,
			&commentUserPhoto.User.Email,
			&commentUserPhoto.Photo.Id,
			&commentUserPhoto.Photo.Title,
			&commentUserPhoto.Photo.Caption,
			&commentUserPhoto.Photo.PhotoUrl,
			&commentUserPhoto.Photo.UserId,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong " + err.Error())
		}

		commentsUserPhoto = append(commentsUserPhoto, commentUserPhoto)
	}

	result := comment_repository.CommentUserPhotoMapped{}
	return result.HandleMappingCommentsUserPhoto(commentsUserPhoto), nil
}

// GetCommentById implements comment_repository.CommentRepository.
func (commentRepo *commentRepositoryImpl) GetCommentById(commentId int) (*comment_repository.CommentUserPhotoMapped, errs.Error) {

	var commentUserPhoto comment_repository.CommentUserPhoto
	rows := commentRepo.db.QueryRow(getCommentByIdQuery, commentId)

	err := rows.Scan(
		&commentUserPhoto.Comment.Id,
		&commentUserPhoto.Comment.UserId,
		&commentUserPhoto.Comment.PhotoId,
		&commentUserPhoto.Comment.Message,
		&commentUserPhoto.Comment.CreatedAt,
		&commentUserPhoto.Comment.UpdatedAt,
		&commentUserPhoto.User.Id,
		&commentUserPhoto.User.Username,
		&commentUserPhoto.User.Email,
		&commentUserPhoto.Photo.Id,
		&commentUserPhoto.Photo.Title,
		&commentUserPhoto.Photo.Caption,
		&commentUserPhoto.Photo.PhotoUrl,
		&commentUserPhoto.Photo.UserId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("comment not found")
		}
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	result := comment_repository.CommentUserPhotoMapped{}
	return result.HandleMappingCommentUserPhoto(commentUserPhoto), nil
}

// DeleteComment implements comment_repository.CommentRepository.
func (commentRepo *commentRepositoryImpl) DeleteComment(commentId int) errs.Error {
	tx, err := commentRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	_, err = tx.Exec(deleteCommentQuery, commentId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return nil
}

// UpdateComment implements comment_repository.CommentRepository.
func (commentRepo *commentRepositoryImpl) UpdateComment(commentId int, commentPayload *entity.Comment) (*dto.UpdateCommentResponse, errs.Error) {

	tx, err := commentRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	row := tx.QueryRow(updateCommentQuery, commentId, commentPayload.Message)

	var comment dto.UpdateCommentResponse
	err = row.Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PhotoId,
		&comment.Message,
		&comment.UpdatedAt,
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
