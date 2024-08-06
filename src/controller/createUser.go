package controller

import (
	"fmt"

	"github.com/Chilengue/go.git/src/configuration/rest_err"
	"github.com/Chilengue/go.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c * gin.Context){

	var userRequest request.UserRequest

	c.ShouldBindJSON(&userRequest)

	if err := c.ShouldBindJSON(&userRequest); err !=nil{
		restErr := rest_err.NewBadRequestError(
		fmt.Sprintf("there are some ibcorrect filds, error=%s\n", err.Error()))

		c.JSON(int(restErr.Code), restErr)
		return
	}
	fmt.Println(userRequest)
}