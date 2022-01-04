package usermodels

import (
	"bookstore_user-api/databaseconnection/postgres"
	"bookstore_user-api/utils/errorparser"
	"bookstore_user-api/utils/errors"
	"fmt"
)

type User struct {
	UserId      int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func New() *User {
	return &User{}
}

func (user *User) Validate() (e *errors.RestError) {
	if user.Email == "" {
		return errors.NewBadRequestError("Email field is mandatory!")
	}
	if user.FirstName == "" {
		return errors.NewBadRequestError("Firstname field is mandatory!")
	}
	if user.LastName == "" {
		return errors.NewBadRequestError("LastName field is mandatory!")
	}
	return nil
}

func (u *User) Save() (err *errors.RestError) {
	insertUserQuery := `INSERT INTO users (first_name, last_name, age, email, date_created) VALUES ($1, $2, $3, $4, $5)`

	dbclient := postgres.Client

	stmt, queryStmtErr := dbclient.Prepare(insertUserQuery)
	if queryStmtErr != nil {
		return errors.NewInternalServerError("Failed to create query statement!")
	}

	defer stmt.Close()

	_, sqlError := stmt.Exec(u.FirstName, u.LastName, u.Age, u.Email, u.DateCreated)
	if sqlError != nil {
		fmt.Println("sqlError while saving user data:", sqlError)
		return errorparser.ParseError(sqlError)
	}

	return nil
}

func (u *User) FetchUser() (err *errors.RestError) {
	getUserQuery := `SELECT id, first_name, last_name, age, email, date_created FROM users WHERE id=$1`

	dbclient := postgres.Client

	stmt, prepareQueryStmtErr := dbclient.Prepare(getUserQuery)
	if prepareQueryStmtErr != nil {
		fmt.Println("Error while preparing statement for fetching user:", prepareQueryStmtErr)
		return errors.NewInternalServerError("Something went wrong while fetching user!")
	}

	defer stmt.Close()

	result := stmt.QueryRow(u.UserId)
	if err := result.Scan(&u.UserId, &u.FirstName, &u.LastName, &u.Age, &u.Email, &u.DateCreated); err != nil {
		fmt.Println("Error while fetching row for user id:", u.UserId, "err:", err)
		return errorparser.ParseError(err)
	}

	return nil
}

func (u *User) Update() (err *errors.RestError) {
	updateUserQuery := `UPDATE users SET first_name=$1, last_name=$2, age=$3, email=$4 WHERE id=$5`

	dbclient := postgres.Client

	stmt, prepareQueryErr := dbclient.Prepare(updateUserQuery)
	if prepareQueryErr != nil {
		fmt.Println("Error while prepare update user query:", prepareQueryErr)
		return errors.NewInternalServerError("Something went wrong while updating user!")
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(u.FirstName, u.LastName, u.Age, u.Email, u.UserId)
	if updateErr != nil {
		fmt.Println("Error while executing update user query:", updateErr)
		return errors.NewInternalServerError("Unexpected error occurred while updating user!")
	}

	return nil
}

func (u *User) Delete() (err *errors.RestError) {
	deleteUserQuery := `DELETE FROM users WHERE id=$1`

	dbclient := postgres.Client

	stmt, prepQueryErr := dbclient.Prepare(deleteUserQuery)
	if prepQueryErr != nil {
		fmt.Println("Error while preparing delete user query:", prepQueryErr)
		return errors.NewInternalServerError("Something went wrong while deleting user!")
	}

	defer stmt.Close()

	_, delErr := stmt.Exec(u.UserId)
	if delErr != nil {
		fmt.Println("Error while executing delete query:", delErr)
		return errors.NewInternalServerError("Unexpected error occurred while deleting user!")
	}

	return nil
}
