package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "signup route "})
}
