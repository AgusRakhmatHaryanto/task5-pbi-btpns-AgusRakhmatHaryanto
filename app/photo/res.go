package photo

import (
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/models"
)


type ResPhoto struct {
	ID        int 		`json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int		`json:"user_id"`
	User      models.User
}

type ResPhotoDefault struct {
	ID        int		`json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
}


func FormatPhoto(photo *models.Photo, typeRes string) interface{} {
	var formatter interface{}

	if typeRes == "default" {
		formatter = ResPhotoDefault{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
		}
	} else {
		formatter = ResPhoto{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserID:   photo.UserID,
			User:     *photo.User,
		}
	}

	return formatter
}
