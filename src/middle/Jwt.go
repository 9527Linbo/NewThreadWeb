package middle

import (
	"NewThread/src/result"
	"NewThread/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if len(auth) == 0 {
		result.CommonResp(c, http.StatusInternalServerError, result.NeedLogin, result.EmptyData)
		return
	}
	// 校验token，只要出错直接拒绝请求
	_, err := utils.ParseToken(auth)
	if err != nil {
		print(err.Error())
		result.CommonResp(c, http.StatusInternalServerError, result.RefusedRequest, result.EmptyData)
		return
	}
}
