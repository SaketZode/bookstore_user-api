package app

import (
	"bookstore_user-api/controllers/pingcontroller"
	"bookstore_user-api/controllers/usercontroller"
)

func mapUrls() {
	pingcontroller := pingcontroller.New()
	usercontroller := usercontroller.New()

	router.GET("/ping", pingcontroller.Ping)
	router.GET("/users/:user_id", usercontroller.GetUserById)
	router.POST("/users", usercontroller.CreateUser)
	router.PUT("/users/:user_id", usercontroller.UpdateUser)
	router.PATCH("/users/:user_id", usercontroller.UpdateUser)
	router.DELETE("/users/:user_id", usercontroller.DeleteUser)
}
