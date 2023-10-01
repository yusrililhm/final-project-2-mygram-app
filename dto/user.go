package dto

type NewUserRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty" example:"monday.day@email.com"`
	Age      uint   `json:"age" valid:"required~Age can't be empty" example:"21"`
	Password string `json:"password" valid:"required~Password can't be empty" example:"secret"`
}

type UserRequest struct {
	Email    string `json:"email" valid:"required~can't be empty" example:"monday.day@email.com"`
	Password string `json:"password" valid:"required~can't be empty" example:"secret"`
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
