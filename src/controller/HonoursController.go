package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HonoursList(c *gin.Context) {
	data, err := logic.NewHonoursService().HonoursList()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
