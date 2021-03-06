package controllers

import (
	"errors"
	"fmt"
	"go_bbs/dao/mysql"
	"go_bbs/logic"
	"go_bbs/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// LoginHandler 用户登录接口
// @Router /api/v1/login [post]
// @Summary 登录接口
// @Accept application/json
// @Produce application/json
// @Param login body _RequestLogin true "登录请求需要上传的json"
// @Success 200 {object} _ResponseLogin

func RegisterHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamRegister)

	fmt.Println(p)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("RegisterHandler with invalid param", zap.Error(err))
		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 业务处理
	err := logic.Register(p)
	if err != nil {
		zap.L().Error("register failed", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.UserAleadyExists) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeInvalidParam)
		}
		return
	}
	// 返回响应
	ResponseSuccess(c, "register success")

}
