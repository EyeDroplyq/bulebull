package controller

import (
	"bulebell/dao/mysql"
	"bulebell/model"
	"bulebell/service"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数，进行参数校验
	p := new(model.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("signUp param is invalid.", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//如果不是validator检验的错误的话
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			//如果是validator检验的错误的话，使用翻译器将错误信息打印出来
			ResponseErrorCustomer(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
	}

	//2.注册
	if err := service.SignUp(p); err != nil {
		zap.L().Error("sigup is failed,err=", zap.Error(err))
		if errors.Is(err, mysql.UserExist) {
			ResponseError(c, CodeUserExist)
			return
		} else {
			ResponseErrorCustomer(c, CodeServerBusy, "注册失败")
			return
		}
	}
	//3.返回结果
	ResponseSuccess(c, nil)
	return
}
func LoginHandler(c *gin.Context) {
	p := new(model.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login param is invalid,err=", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		} else {
			ResponseErrorCustomer(c, CodeServerBusy, removeTopStruct(errs.Translate(trans)))
			return
		}
	}
	if err := service.Login(p); err != nil {
		zap.L().Error("login is failed,err=", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, nil)
	return
}
