package errorparser

import (
	"bookstore_user-api/utils/errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func ParseError(err error) *errors.RestError {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		fmt.Println("Error while parsing into postgres error struct")
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError("Record with given ID not found!")
		}
		return errors.NewInternalServerError("Something went wrong!")
	}
	switch pqErr.Code {
	case "23505":
		return errors.NewBadRequestError("Duplicate entry found!")
	default:
		fmt.Println("Unexpected SQL error:", pqErr.Code, "Detail:", pqErr.Detail)
		return errors.NewInternalServerError("Unexpected error occurred!")
	}
}
