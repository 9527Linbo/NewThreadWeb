package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"NewThread/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersIcon(c *gin.Context) {

	ids, vis := c.GetQueryArray("userids")
	if !vis {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	data, err := logic.NewUserService().UserIcon(utils.StringToInt(ids))
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
