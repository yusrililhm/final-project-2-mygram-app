package dto

import (
	"time"
)

type NewCommentRequest struct {
	PhotoId int    `json:"photo_id" example:"1"`
	Message string `json:"message" valid:"required~Message can't be empty" example:"so beautifull"`
}

type NewCommentResponse struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" valid:"required~Message can't be empty" example:"so beautifull"`
}
