package controllers

import (
	"net/http"
	"encoding/json"

	resUser "github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/app/user"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/helpers"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}


func (h *UserController) Register(c *gin.Context) {
    var user models.User
	c.ShouldBindJSON(&user)
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		res := helpers.ApiResponseFormatter("Register account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	user.Password = hashedPassword


    if err := h.DB.Debug().Create(&user).Error; err != nil {
        err = helpers.FormatError(err)
        errmessage := gin.H{"errors": err}
        res := helpers.ApiResponseFormatter("Register account failed", http.StatusUnprocessableEntity, "error", errmessage)
        c.JSON(http.StatusUnprocessableEntity, res)
        return
    }

    formatter := resUser.FormatUserResponse(user, "")
    res := helpers.ApiResponseFormatter("Account has been registered", http.StatusOK, "success", formatter)
    c.JSON(http.StatusOK, res)
}



func (h *UserController) Login(c *gin.Context) {

	var user models.User
	c.ShouldBindJSON(&user)

	inputPassword := user.Password
	err := h.DB.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		res := helpers.ApiResponseFormatter("Login failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	comparePassword := helpers.CheckPasswordHash(user.Password, inputPassword)
	if !comparePassword {
		res := helpers.ApiResponseFormatter("Login failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		res := helpers.ApiResponseFormatter("Login failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	formatter := resUser.FormatUserResponse(user, token)
	res := helpers.ApiResponseFormatter("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *UserController) Update(c *gin.Context) {
	var oldUser models.User
	var newUser models.User

	id := c.Param("userId")

	err := h.DB.First(&oldUser, id).Error
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helpers.ApiResponseFormatter(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		response := helpers.ApiResponseFormatter(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.DB.Model(&oldUser).Updates(newUser).Error
	if err != nil {
		response := helpers.ApiResponseFormatter(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponseFormatter(http.StatusOK, "success", nil, "User Updated Succesfully")
	c.JSON(http.StatusOK, response)
}

func (h *UserController) Delete(c *gin.Context) {
	var user models.User

	id := c.Param("userId")
	err := h.DB.First(&user, id).Error
	if err != nil {
		response := helpers.ApiResponseFormatter(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.DB.Delete(&user).Error
	if err != nil {
		response := helpers.ApiResponseFormatter(http.StatusUnprocessableEntity, "error", nil, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponseFormatter(http.StatusOK, "success", nil, "User Deleted Succesfully")
	c.JSON(http.StatusOK, response)
}
