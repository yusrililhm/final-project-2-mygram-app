package photo_repository

import (
	"myGram/entity"
	"time"
)

type PhotoUser struct {
	Photo entity.Photo
	User  entity.User
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoUserMapped struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

func (pum *PhotoUserMapped) HandleMappingPhotoWithUser(photoUser []PhotoUser) []PhotoUserMapped {
	photosUserMapped := []PhotoUserMapped{}

	for _, eachPhotoUser := range photoUser {
		photoUserMapped := PhotoUserMapped{
			Id:        eachPhotoUser.Photo.Id,
			Title:     eachPhotoUser.Photo.Title,
			Caption:   eachPhotoUser.Photo.Caption,
			PhotoUrl:  eachPhotoUser.Photo.PhotoUrl,
			UserId:    eachPhotoUser.Photo.UserId,
			CreatedAt: eachPhotoUser.Photo.CreatedAt,
			UpdatedAt: eachPhotoUser.Photo.UpdatedAt,
			User: User{
				Email:    eachPhotoUser.User.Email,
				Username: eachPhotoUser.User.Username,
			},
		}

		photosUserMapped = append(photosUserMapped, photoUserMapped)
	}

	return photosUserMapped
}
