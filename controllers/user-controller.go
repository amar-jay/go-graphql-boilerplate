package controllers

import (
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
	"github.com/gin-gonic/gin"
)

/**
* ---- Input Types -----
*/

type UserController interface {
  Register(*gin.Context)
  Update(*gin.Context)
  Login(*gin.Context)
  GetUserByID(*gin.Context)
  GetProfile(*gin.Context)
  ResetPassword(*gin.Context)
  ForgotPassword(*gin.Context)
}

type userController struct {
  us userservice.UserService
  as authservice.AuthService
  es emailservice.EmailService
}


// NewUserService creates a an instance of User Service 
func NewUserController(us userservice.UserService, as authservice.AuthService, es emailservice.EmailService) UserController {
	return &userController{us, as, es}
}

/**
* ----- Routes -----
*/

func (user *userController) Login(ctx *gin.Context){

}

func (user *userController) Register(ctx *gin.Context) {
}

func (user *userController) Update(ctx *gin.Context) {
}

func (user *userController) ResetPassword(ctx *gin.Context) {
}

func (user *userController) ForgotPassword(ctx *gin.Context) {
}

func (user *userController) GetProfile(ctx *gin.Context) {
}

func (user *userController) GetUserByID(ctx *gin.Context) {
}
