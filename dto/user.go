package dto

type NewUserRequest struct {
	Username string `json:"username" validate:"required, unique" example:"anyujin"`
	Email    string `json:"email" validate:"required, unique" example:"yujin.an@email.com"`
	Password string `json:"password" validate:"required, min=6" example:"secret"`
	Age      uint   `json:"age" validate:"required, gte=8" example:"20"`
}
