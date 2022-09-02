package controller

import (
	"github.com/wisnuragaprawida/golang-rest-api/delivery/api"
	"github.com/wisnuragaprawida/golang-rest-api/model"
	"github.com/wisnuragaprawida/golang-rest-api/usecase"
	"github.com/wisnuragaprawida/golang-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	ucUser usecase.LoginUserUseCase
	api.BaseApi
}

func (u *UserController) loginUser(c *gin.Context) {
	var newUser model.User
	err := u.ParseRequestBody(c, &newUser)
	if err != nil {
		u.Failed(c, utils.RequiredError())
		return
	}
	_, err = u.ucUser.LoginUser(&newUser)
	if err != nil {
		u.Failed(c, err)
		return
	}
	u.Success(c, newUser)
}

func NewUserController(
	router *gin.Engine,
	ucUser usecase.LoginUserUseCase) *UserController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := UserController{
		router: router,
		ucUser: ucUser,
	}

	// ini method-method nya
	router.POST("/user/login", controller.loginUser)
	return &controller
}
