package userservice

import (
	"bookstore_user-api/models/usermodels"
	"bookstore_user-api/utils/errors"
)

type UserService struct {
}

func New() *UserService {
	return &UserService{}
}

func (userservice UserService) CreateUser(user *usermodels.User) (*usermodels.User, *errors.RestError) {
	validationErr := user.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	saveError := user.Save()
	if saveError != nil {
		return nil, saveError
	}
	return user, nil
}

func (userservice UserService) GetAllUsers() (userlist []usermodels.User, err *errors.RestError) {
	return userlist, nil
}

func (userservice UserService) GetUser(id int) (user *usermodels.User, err *errors.RestError) {
	user = &usermodels.User{UserId: id}
	fetchErr := user.FetchUser()
	if fetchErr != nil {
		return nil, fetchErr
	}
	return user, nil
}

func (userservice UserService) UpdateUser(isPartialUpdate bool, newuser usermodels.User) (user *usermodels.User, err *errors.RestError) {
	currentuser, getErr := userservice.GetUser(newuser.UserId)
	if getErr != nil {
		return nil, getErr
	}

	if isPartialUpdate {

	} else {
		currentuser.FirstName = newuser.FirstName
		currentuser.LastName = newuser.LastName
		currentuser.Email = newuser.Email
		currentuser.Age = newuser.Age
	}
	updateErr := currentuser.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return currentuser, nil
}
