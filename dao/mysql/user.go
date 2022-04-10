package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go_bbs/models"

	"go.uber.org/zap"
)

const serect = "tracydzf"

// 定义 error的常量方便判断
var (
	UserAleadyExists = errors.New("用户已存在")
	WrongPassword    = errors.New("密码不正确")
	UserNoExists     = errors.New("用户不存在")
)

func InsertUser(user *models.User) error {
	// 密码要加密保存
	user.Password = encryptPassword(user.Password)
	sqlstr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err := db.Exec(sqlstr, user.UserId, user.Username, user.Password)
	if err != nil {
		zap.L().Error("InsertUser dn error", zap.Error(err))
		return err
	}
	return nil
}

func CheckUserExist(username string) error {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlstr, username)

	if err != nil {
		zap.L().Error("CheckUserExist dn error", zap.Error(err))
		return err
	}
	if count > 0 {
		return UserAleadyExists
	}
	return nil
}

// 加密密码
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(serect))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
