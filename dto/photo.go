package dto

import "time"

type NewPhotoRequest struct {
	Title    string `json:"title" valid:"required~Title can't be empty" example:"monday awesome"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL can't be empty" example:"https://www.pinterest.com/pin/807973989398829161/"`
	Caption  string `json:"caption" example:"Hello I'm Monday from Weeekly, hopefully You can do this!"`
}

type PhotoResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
