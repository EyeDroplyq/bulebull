package controller

import (
	"bulebell/model"
	"bulebell/service"
	"net/http"

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
			c.JSON(http.StatusOK, gin.H{
				"msg": "参数校验失败",
			})
			return
		} else {
			//如果是validator检验的错误的话，使用翻译器将错误信息打印出来
			c.JSON(http.StatusOK, removeTopStruct(errs.Translate(trans)))
			return
		}
	}

	//2.注册
	if err := service.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	//3.返回结果
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
