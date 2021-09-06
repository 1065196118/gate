package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/social-mediam-users/controller"
	"github.com/social-mediam-users/repositories"
)

var (
	userRepository repositories.UserRepository = repositories.NewUserRepository()
	userController controller.UserController   = controller.NewUserController(userRepository)
)

func SetupUserRoute(route *gin.Engine) {
	api := route.Group("/users")
	{
		api.POST("/me", userController.Me)
		api.POST("/register", userController.Store)
		api.POST("/login", userController.Login)
	}
}
