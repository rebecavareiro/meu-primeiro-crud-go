package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rebecavarreiro/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/rebecavarreiro/meu-primeiro-crud-go/src/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(

			fmt.Sprintf("Esxite alguns campos incorretos, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
