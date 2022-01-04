package usercontroller

import (
	"net/http"
	"strconv"

	"bookstore_user-api/models/usermodels"
	"bookstore_user-api/services/userservice"
	"bookstore_user-api/utils/errors"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func New() *UserController {
	return &UserController{}
}

var userservices = userservice.New()

func (usercontroller UserController) CreateUser(c *gin.Context) {
	var user = usermodels.New()
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid data received!!")
		return
	}

	result, userCreationErr := userservices.CreateUser(user)
	if userCreationErr != nil {
		c.JSON(userCreationErr.StatusCode, userCreationErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (usercontroller UserController) GetUsers(c *gin.Context) {
	result, uErr := userservices.GetAllUsers()
	if uErr != nil {
		c.JSON(uErr.StatusCode, uErr)
		return
	}
	c.JSON(http.StatusAccepted, result)
}

func (usercontroller UserController) GetUserById(c *gin.Context) {
	userid := c.Param("user_id")
	id, err := strconv.ParseInt(userid, 10, 32)
	if err != nil {
		perror := errors.NewBadRequestError("please pass valid user id :-|")
		c.JSON(perror.StatusCode, perror)
		return
	}
	user, fetchError := userservices.GetUser(int(id))
	if fetchError != nil {
		c.JSON(fetchError.StatusCode, fetchError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (usercontroller UserController) UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "UpdateUser Functionality not implemented!!")
}

func (usercontroller UserController) DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "DeleteUser Functionality not implemented!!")
}
