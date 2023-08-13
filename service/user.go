package service

import (
	"bulebell/dao/mysql"
	"bulebell/model"
	"bulebell/pkg/snowflake"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secrt = "lyq@163.com"

func SignUp(p *model.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err = mysql.CheckExistUser(p.Username); err != nil {
		return errors.New("用户存在")
	}
	userID := snowflake.GenID()
	//2.密码进行加密
	p.Password = encryptPassword(p.Password)
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
		Gender:   p.Gender,
		UserID:   userID,
	}
	//3.插入数据库
	if err = mysql.InsertUser(user); err != nil {
		return errors.New("保存用户失败")
	}
	return nil
}

// encryptPassword:对明文密码进行md5加密
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secrt))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
