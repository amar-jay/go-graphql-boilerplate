package controllers

import (
	"fmt"
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
  Email string `json:"email"`
  Password string `json:"password"`
}

type UserOutput struct {
  ID uint `json:"ID"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  Role string `json:"role"`
  Active bool `json:"acive"`

}

type ErrorOutput struct {
  Msg string `json:"message"`
  Summary string `json:"summary"`
  Code int `json:"statusCode"`
}
/**
* ---- Input Types -----
*/

type UserController interface {
  Register(*gin.Context)
  Update(*gin.Context)
  login(ctx *gin.Context, u *user.User) error
  Login(ctx *gin.Context)
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

func (user *userController) Login(ctx *gin.Context) {
}

func (userctrl *userController) Register(ctx *gin.Context) {
  // TODO: read the user input
  var userInput UserInput
  if err := ctx.ShouldBindJSON(&userInput); err != nil {
    errSum := ErrorOutput{
      Msg: err.Error(),
      Summary: "Userinput Error",
      Code: http.StatusBadRequest,
    }

    fmt.Println(errSum)
    HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
    return
  } 

  u := userctrl.inputToUser(userInput)
  // TODO: create a user

  if err := userctrl.us.Register(&u); err != nil {
    
    ctx.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  //err := fmt.Errorf("Not implemented")
  //ctx.AbortWithError(http.StatusBadRequest, err)
  // TODO: send a welcome message
  // TODO: login the user
  if err := userctrl.login(ctx, &u); err != nil {
    ctx.JSON(http.StatusInternalServerError, err.Error())
    return
  }
}
func (user *userController) Update(ctx *gin.Context) {
  // TODO: read the user input
  // TODO: get the user from the database 
  // TODO: check password
  // TODO: login the user 
    fmt.Println("🔎 Check out the user service")
}
func (user *userController) ResetPassword(ctx *gin.Context) {
    fmt.Println("🔎 Check out the user controller")
}

func (user *userController) ForgotPassword(ctx *gin.Context) {
    fmt.Println("🔎 Check out the user controller")
}

func (user *userController) GetProfile(ctx *gin.Context) {
    fmt.Println("🔎 Check out the user controller")
}

func (user *userController) GetUserByID(ctx *gin.Context) {
    fmt.Println("🔎 Check out the user controller")
}

/**
* --- Other methods
*/

func (userctrl *userController) inputToUser(input UserInput) user.User {
  return user.User{
    Email:  input.Email,
    Password: input.Password,
  }
}


func (userctrl *userController) mapToUserOutput(input *user.User) *UserOutput {
  return &UserOutput{
    ID: input.ID,
    Email: input.Email,
    FirstName: input.FirstName,
    LastName: input.LastName,
    Active: input.Active,
    Role: input.Role,
  }
}
 func (userctrl *userController) login(ctx *gin.Context, u *user.User) error {
   token, err := userctrl.as.IssueToken(*u)
   if err != nil {
     return err
   }
   userOutput := userctrl.mapToUserOutput(u)
   out := gin.H{"token": token, "user": userOutput}
   HttpResponse(ctx, http.StatusOK, "ok", out)
   return nil
 }
