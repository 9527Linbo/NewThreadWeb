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

func HonoursStudents(c *gin.Context) {
	data, err := logic.NewHonoursService().HonoursStudents()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func HonoursProjects(c *gin.Context) {
	data, err := logic.NewHonoursService().HonoursProjects()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func HonoursMilestone(c *gin.Context) {
	data, err := logic.NewHonoursService().HonoursMilestone()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func HonoursMilestones(c *gin.Context) {
	data, err := logic.NewHonoursService().HonoursMilestones()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
