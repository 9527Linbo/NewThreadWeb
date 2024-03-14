package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadShare(c *gin.Context) {

	data, err := logic.NewPostReadService().ReadShareInfo(c)

	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func ReadNews(c *gin.Context) {

	data, err := logic.NewPostReadService().ReadNewsInfo(c)

	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func ReadActivity(c *gin.Context) {

	data, err := logic.NewPostReadService().ReadActivityInfo(c)

	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	result.CommonResp(c, http.StatusOK, result.Success, data)
}
