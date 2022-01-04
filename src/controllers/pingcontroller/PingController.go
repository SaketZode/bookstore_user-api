package pingcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
}

func New() *PingController {
	return &PingController{}
}

func (pingcontroller *PingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "bookstore_user-api service is running :-)")
}
