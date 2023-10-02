package dto

import "time"

type NewUserRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"monday.day@email.com"`
	Age      uint   `json:"age" valid:"required~Age can't be empty, range(8|150)~Minimum age is 8" example:"21"`
	Password string `json:"password" valid:"required~Password can't be empty, length(6|255)~Minimum password is 6 length" example:"secret"`
}

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"monday.day@email.com"`
	Password string `json:"password" valid:"required~Password can't be empty" example:"secret"`
}

type UserUpdateRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"monday.day@email.com"`
}

type UserUpdateResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
