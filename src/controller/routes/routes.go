package routes

import (
	"github.com/Chilengue/go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
r.GET("/getUserById/:userId", controller.FingUserByID)
r.GET("/getUserByEmail/:userEmail", controller.FingUserByEmail)
r.POST("/createUser", controller.CreateUser)
r.PUT("/updateUser/:userId", controller.UpdateUser)
r.DELETE("/deleteUser/:userId", controller.DelteUser)
}
