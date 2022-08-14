package controllers

import (
	"net/http"

	"github.com/amar-jay/go-api-boilerplate/domain/user"
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
	"github.com/gin-gonic/gin"
)

/**
*  --- Types ---
 */
type UserInput struct {
  email string `json:"email"`
  password string `json:"password"`
}

type userOutput struct {
  ID string `json:"ID"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  Role string `json:"role"`
  Active string `json:"acive"`

}
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

func (userctrl *userController) Register(ctx *gin.Context) {
  // TODO: read the user input
  var userInput UserInput
  if err := ctx.ShouldBind(userInput); err != nil {
    ctx.JSON(http.StatusBadRequest, err.Error())
    return
  } 

  u := userctrl.inputToUser(userInput)
  // TODO: create a user

  if err := userctrl.us.Register(&u); err != nil {
    ctx.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  // TODO: send a welcome message
  // TODO: login the user
}
func (user *userController) Update(ctx *gin.Context) {
  // TODO: read the user input
  // TODO: get the user from the database 
  // TODO: check password
  // TODO: login the user 
}
func (user *userController) ResetPassword(ctx *gin.Context) {
}

func (user *userController) ForgotPassword(ctx *gin.Context) {
}

func (user *userController) GetProfile(ctx *gin.Context) {
}

func (user *userController) GetUserByID(ctx *gin.Context) {
}

/**
* --- Other methods
*/

func (userctrl *userController) inputToUser(input UserInput) *user.User {
  return &user.User{
    Email:  input.email,
    Password: input.password,
  }
}

