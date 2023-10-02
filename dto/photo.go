package dto

type NewPhotoRequest struct {
	Title    string `json:"title" valid:"required~Title can't be empty" example:"monday awesome"`
	PhotoUrl string `json:"caption" valid:"required~Photo URL can't be empty" example:"https://www.pinterest.com/pin/807973989398829161/"`
	Caption  string `json:"photo_url" example:"Hello I'm Monday from Weeekly, hopefully You can do this!"`
}
