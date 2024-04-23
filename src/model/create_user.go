package model

import (
	"fmt"

	"github.com/rebecavarreiro/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/rebecavarreiro/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("jouney", "createUser"))
	// email nviar com codigo

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
