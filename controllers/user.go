package controllers

import (
	"encoding/json"
	"fmt"
	userRes "github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/app/user"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/helpers"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

func (h *userController) Register(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user.Password = hashedPassword

	err = h.db.Debug().Create(&user).Error
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userRes.FormatUserResponse(user, "")
	response := helpers.ApiResponseFormatter("success", http.StatusOK, "User Registered Succesfully", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var user models.User

	c.ShouldBindJSON(&user)

	Inputpassword := user.Password
	err := h.db.Debug().Where("email = ?", user.Email).Find(&user).Error
	if err != nil {
		fmt.Println("Error querying database:", err)
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, "Login Failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	comparePass := helpers.CheckPasswordHash(user.Password, Inputpassword)

	fmt.Println("Password Comparison Result:", comparePass)
	fmt.Println("Stored Hash Password:", user.Password)
	fmt.Println("Stored Hash Passwordiii:", Inputpassword)
	fmt.Println("Input Hash Password:", helpers.CheckPasswordHash(user.Password, Inputpassword))

	if !comparePass {
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, "Login Failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println("Error generating token:", err)
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, "Login Failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := userRes.FormatUserResponse(user, token)
	response := helpers.ApiResponseFormatter("success", http.StatusOK, "User Login Succesfully", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) Update(c *gin.Context) {
	var oldUser models.User
	var newUser models.User

	id := c.Param("userId")

	err := h.db.First(&oldUser, id).Error
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.db.Model(&oldUser).Updates(newUser).Error
	if err != nil {
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponseFormatter("success", http.StatusOK, "User Updated Succesfully", nil)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Delete(c *gin.Context) {
	var user models.User

	id := c.Param("userId")
	err := h.db.First(&user, id).Error
	if err != nil {
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.db.Delete(&user).Error
	if err != nil {
		response := helpers.ApiResponseFormatter("error", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.ApiResponseFormatter("success", http.StatusOK, "User Deleted Succesfully", nil)
	c.JSON(http.StatusOK, response)
}
