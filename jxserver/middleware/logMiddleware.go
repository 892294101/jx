// 操作日志中间件
// author xiaoRui

package middleware

import (
	"github.com/jx/jxserver/api/dao"
	"github.com/jx/jxserver/api/entity"
	"github.com/jx/jxserver/common/util"
	"github.com/jx/jxserver/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		sysAdmin, _ := jwt.GetAdmin(c)
		if method != "get" {
			log := entity.SysOperationLog{
				AdminId:    sysAdmin.ID,
				Username:   sysAdmin.Username,
				Method:     method,
				Ip:         c.ClientIP(),
				Url:        c.Request.URL.Path,
				CreateTime: util.HTime{Time: time.Now()},
			}
			dao.CreateSysOperationLog(log)
		}
	}
}
