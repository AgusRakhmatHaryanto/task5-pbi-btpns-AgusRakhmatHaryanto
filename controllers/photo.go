package controllers

import (
	"net/http"

	"time"

	resPhoto "github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/app/photo"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/helpers"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PhotoController struct {
	DB *gorm.DB
}

func NewPhotoController(DB *gorm.DB) PhotoController {
	return PhotoController{DB}
}

func (s *PhotoController) GetPhotos(c *gin.Context) {
	var  userPhotos models.Photo
	err := s.DB.Preload("User").Find(&userPhotos).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userPhotos.ID == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	formatter := resPhoto.FormatPhoto(&userPhotos, "default")
	response := helpers.ApiResponseFormatter("List of Photos", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (s *PhotoController) CreatePhoto(c *gin.Context) {

	var inputPhoto models.Photo
	var photoCount int64
	currentUser := c.MustGet("currentUser").(models.User)

	
	s.DB.Model(&inputPhoto).Where("user_id = ?", currentUser.ID).Count(&photoCount)
	if photoCount>0 {
		data := gin.H{
			"error": "you can only upload one photo",
			"is_uploaded": false,
		}
		response := helpers.ApiResponseFormatter("Failed to create photo", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.ApiResponseFormatter("Create photo failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("photo_url")
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.ApiResponseFormatter("Create photo failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ext := file.Filename
	path := "uploads/" + uuid.New().String() + ext

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		
		response := helpers.ApiResponseFormatter("Create photo failed", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	s.AddPhoto(input, path, currentUser.ID)

	data := gin.H{
		"is_uploaded": true,
	}

	response := helpers.ApiResponseFormatter("Photo successfuly created", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
	return
}

func (s *PhotoController) AddPhoto(userPhoto models.Photo, file string, userID uuid.UUID) error {
	savePhoto := models.Photo{
		UserID:    userID,
		Title:     userPhoto.Title,
		Caption:   userPhoto.Caption,
		PhotoUrl:  file,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.DB.Debug().Create(&savePhoto).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PhotoController) UpdatePhoto(c *gin.Context) {

	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	err := s.DB.Where("user_id = ?", currentUser.ID).First(&userPhoto).Error
	if err != nil {
		res := helpers.ApiResponseFormatter("Failed to update photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var inputPhoto models.Photo

	err = c.ShouldBindJSON(&inputPhoto)
	if err != nil {
		res := helpers.ApiResponseFormatter("Failed to update photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	file, err := c.FormFile("photo_url")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helpers.ApiResponseFormatter("Failed to update photo", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	ext := file.Filename
	path := "uploads/" + uuid.New().String() + ext

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		res := helpers.ApiResponseFormatter("Failed to update photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	s.UpdateNewPhoto(inputPhoto, &userPhoto, path)

	data := resPhoto.FormatPhoto(&userPhoto, "default")
	res := helpers.ApiResponseFormatter("Photo successfuly updated", http.StatusOK, "success", data)
	
	c.JSON(http.StatusOK, res)
	}

	
	func (h *PhotoController) UpdateNewPhoto(oldPhoto models.Photo, newPhoto *models.Photo, path string) error {
		newPhoto.Title = oldPhoto.Title
		newPhoto.Caption = oldPhoto.Caption
		newPhoto.PhotoUrl = path
	
		err := h.DB.Save(&newPhoto).Error
		if err != nil {
			return err
		}
	
		return nil
	}

	func (s *PhotoController) DeletePhoto(c *gin.Context) {
		var userPhoto models.Photo
		currentUser := c.MustGet("currentUser").(models.User)

		err := s.DB.Where("user_id = ?", currentUser.ID).First(&userPhoto).Error
		if err != nil {
			res := helpers.ApiResponseFormatter("Failed to delete photo", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		data := gin.H{
			"is_deleted": true,
		}

		response := helpers.ApiResponseFormatter("Photo successfuly deleted", http.StatusOK, "success", data)

		c.JSON(http.StatusOK, response)
		}
	