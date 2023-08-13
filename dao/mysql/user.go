package mysql

import (
	"bulebell/model"
	"errors"

	"go.uber.org/zap"
)

// CheckExistUser:根据用户名判断用户是否已经存在
func CheckExistUser(Username string) error {
	sqlstr := `select count(*) from user where username=?`
	var count int
	err := db.Get(&count, sqlstr, Username)
	if err != nil {
		zap.L().Error("查询用户是否存在错误", zap.Error(err))
		return errors.New("查询错误")
	}
	return nil
}
func InsertUser(user *model.User) (err error) {
	sqlstr := `insert into user(user_id,username,password,gender) values (?,?,?,?)`
	_, err = db.Exec(sqlstr, user.UserID, user.Username, user.Password, user.Gender)
	if err != nil {
		zap.L().Error("保存用户失败", zap.Error(err))
		return err
	}
	return nil
}
