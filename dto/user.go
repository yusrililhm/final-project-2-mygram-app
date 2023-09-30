package dto

type NewUserRequest struct {
	Username string `json:"username" validate:"required, unique" example:"anyujin"`
	Email    string `json:"email" validate:"required, unique" example:"yujin.an@email.com"`
	Age      uint   `json:"age" validate:"required, gte=8" example:"20"`
	Password string `json:"password" validate:"required, min=6" example:"secret"`
}

type UserRequest struct {
	Email    string `json:"email" validate:"required, unique" example:"yujin.an@email.com"`
	Password string `json:"password" validate:"required, min=6" example:"secret"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type GetUserResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
