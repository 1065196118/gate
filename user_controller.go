package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/social-mediam-users/dto"
	"github.com/social-mediam-users/helpers"
	"github.com/social-mediam-users/repositories"
)

type UserController interface {
	Store(c *gin.Context)
	Login(c *gin.Context)
	Me(c *gin.Context)
}

type userController struct {
	repository repositories.UserRepository
}

func NewUserController(repository repositories.UserRepository) UserController {
	return &userController{
		repository: repository,
	}
}

func (repository userController) Store(c *gin.Context) {

	var userJson dto.UserCreateDTO

	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": helpers.Validation(err)})
		return
	}

	isDublicateEmail := repository.repository.IsDublicateEmail(userJson.Email)

	if isDublicateEmail.RowsAffected > 0 {
		helpers.ValidationReturnErrorResponse(c, "email", "unique", "Email is not available")
	}

	isDublicateUsername := repository.repository.IsDublicateUsername(userJson.Username)

	if isDublicateUsername.RowsAffected > 0 {
		helpers.ValidationReturnErrorResponse(c, "username", "unique", "Username is not available")
	}

	isDublicatePhone := repository.repository.IsDublicatePhone(userJson.Phone)

	if isDublicatePhone.RowsAffected > 0 {
		helpers.ValidationReturnErrorResponse(c, "phone", "unique", "Phone is not available")
	}

	repository.repository.Store(userJson)

	helpers.EmailUsersWelcomeMessage(userJson.Email)

	c.JSON(202, gin.H{
		"payload": userJson,
	})
}

func (repository userController) Login(c *gin.Context) {
	var loginDTO dto.UserLoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": helpers.Validation(err)})
		return
	}

	userId := repository.repository.Login(loginDTO)

	if userId != 0 {
		guid := xid.New()
		randomToken := guid.String()
		agent := c.Request.Header.Get("User-Agent")
		clientIP := c.ClientIP()
		if repository.repository.IsLoginFromOtherDevice(userId, agent, clientIP, randomToken) {
			fmt.Println("send Email")
		}
		c.JSON(http.StatusOK, gin.H{
			"payload": randomToken,
		})
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"payload": "Unauthorized",
	})
}

func (repository userController) Me(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Request.Header.Get("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"payload": "",
		})
	}
	user := repository.repository.Me(userId)
	c.JSON(http.StatusOK, gin.H{
		"payload": user,
	})
}
