package routes

import (
	"github.com/gin-gonic/gin"
	"web-service/controllers/user"
)

func SetupRouter(r *gin.Engine) {
	apiGroup := r.Group("/api/v1")

	apiGroup.POST("/signin", user.SigninController)
	apiGroup.POST("/signup", user.SignupController)

}
