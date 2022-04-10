package logic

import (
	"go_bbs/dao/mysql"
	"go_bbs/models"
	"go_bbs/pkg/snowflake"
)

func Register(register *models.ParamRegister) (err error) {
	// 判断用户是否存在
	err = mysql.CheckUserExist(register.UserName)
	if err != nil {
		// db 出错
		return err
	}
	// 生成userid
	userId := snowflake.GenId()

	// 构造一个User db对象

	user := models.User{
		UserId:   userId,
		Username: register.UserName,
		Password: register.Password,
	}
	// 保存数据库
	// 保存数据库
	err = mysql.InsertUser(&user)
	if err != nil {
		return err
	}
	return

}
