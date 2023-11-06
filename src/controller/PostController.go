package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostInfo(c *gin.Context) {
	data, err := logic.NewPostService().PostInfo()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func PostShareInfo(c *gin.Context) {
	data, err := logic.NewPostService().PostShareInfo()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func PostNewsInfo(c *gin.Context) {
	data, err := logic.NewPostService().PostNewsInfo()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func PostActivityInfo(c *gin.Context) {
	data, err := logic.NewPostService().PostActivityInfo()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
