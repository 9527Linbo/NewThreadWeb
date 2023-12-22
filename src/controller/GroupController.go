package controller

import (
	"NewThread/src/logic"
	"NewThread/src/pojo"
	"NewThread/src/result"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GroupInfo(c *gin.Context) {
	data, err := logic.NewGroupService().GroupInfo()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func GroupTeacherInfo(c *gin.Context) {
	data, err := logic.NewGroupService().GroupTeacherInfo()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func GroupStudentInfo(c *gin.Context) {
	var req pojo.Year
	if err := c.ShouldBind(&req); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	data, err := logic.NewGroupService().GroupStudentInfo(req.Year)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func Yearlist(c *gin.Context) {
	data, err := logic.NewGroupService().Yearlist()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
