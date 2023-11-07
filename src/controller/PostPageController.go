package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo
// func PostInfo(c *gin.Context) {
// 	data, err := logic.NewPostPageService().PostInfo()
// 	if err != nil {
// 		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
// 		return
// 	}
// 	result.CommonResp(c, http.StatusOK, result.Success, data)
// }

func PageShareInfo(c *gin.Context) {
	data, err := logic.NewPostPageService().PageShareInfo(c)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func PageNewsInfo(c *gin.Context) {
	data, err := logic.NewPostPageService().PageNewsInfo(c)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func PageActivityInfo(c *gin.Context) {
	data, err := logic.NewPostPageService().PageActivityInfo(c)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
