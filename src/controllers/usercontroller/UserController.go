package usercontroller

import (
	"fmt"
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

func parseUserID(param string) (int, *errors.RestError) {
	id, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		fmt.Println("Error while parsing user ID:", err)
		return 0, errors.NewBadRequestError("Please provide valid user ID :-|")
	}

	return int(id), nil
}

func (usercontroller *UserController) CreateUser(c *gin.Context) {
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

func (usercontroller *UserController) GetUsers(c *gin.Context) {
	result, uErr := userservices.GetAllUsers()

	if uErr != nil {
		c.JSON(uErr.StatusCode, uErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (usercontroller *UserController) GetUserById(c *gin.Context) {
	userId, parseErr := parseUserID(c.Param("user_id"))

	if parseErr != nil {
		c.JSON(parseErr.StatusCode, parseErr)
		return
	}

	user, fetchError := userservices.GetUser(userId)
	if fetchError != nil {
		c.JSON(fetchError.StatusCode, fetchError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (usercontroller *UserController) UpdateUser(c *gin.Context) {
	userId, parseErr := parseUserID(c.Param("user_id"))

	if parseErr != nil {
		c.JSON(parseErr.StatusCode, parseErr)
		return
	}

	var isPartialUpdate bool = false
	if c.Request.Method == http.MethodPatch {
		isPartialUpdate = true
	}

	var newuser = usermodels.New()
	jsonerr := c.ShouldBindJSON(&newuser)
	if jsonerr != nil {
		invalidjsonerr := errors.NewBadRequestError("Invalid data received!")
		c.JSON(invalidjsonerr.StatusCode, invalidjsonerr)
		return
	}

	newuser.UserId = int(userId)
	res, updateErr := userservices.UpdateUser(isPartialUpdate, *newuser)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (usercontroller *UserController) DeleteUser(c *gin.Context) {
	userID, parseError := parseUserID(c.Param("user_id"))
	if parseError != nil {
		c.JSON(parseError.StatusCode, parseError)
		return
	}
	delErr := userservices.DeleteUser(userID)
	if delErr != nil {
		c.JSON(delErr.StatusCode, delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
