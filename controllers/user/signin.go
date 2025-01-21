package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SigninController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "signin route"})
}
