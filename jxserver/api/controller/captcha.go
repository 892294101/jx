// 验证码 控制层
// author xiaoRui

package controller

import (
	"github.com/jx/jxserver/common/result"
	"github.com/gin-gonic/gin"
	"github.com/jx/jxserver/api/service"
)

// @Summary 验证码接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptMake()
	result.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}
