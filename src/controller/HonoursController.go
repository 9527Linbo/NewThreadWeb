package controller

import (
	"NewThread/src/logic"
	"NewThread/src/pojo"
	"NewThread/src/result"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func HonoursList(c *gin.Context) {

	data, err := logic.NewHonoursService().HonoursList()

	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func HonoursGraduate(c *gin.Context) {

	data, err := logic.NewHonoursService().HonoursGraduate()

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

func AddMilestone(c *gin.Context) {

	var vis byte
	vis = 1
	if strings.Contains(c.Request.URL.Path, `/add/project`) {
		vis = 0
	}

	var projectmsg pojo.Project
	//获取参数
	if err := c.BindJSON(&projectmsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	data, err := logic.NewHonoursService().AddMilestone(projectmsg, vis)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, err.Error())
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func AddHonour(c *gin.Context) {
	var honourmsg pojo.Honours
	var OtherMsg struct {
		Type    string `json:"Type"`
		EndTime string `json:"EndTime"`
	}
	//获取参数c.ShouldBindBodyWith(&req,binding.JSON)
	if err := c.ShouldBindBodyWith(&honourmsg, binding.JSON); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	if err := c.ShouldBindBodyWith(&OtherMsg, binding.JSON); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, err.Error())
		return
	}

	data, err := logic.NewHonoursService().AddHonour(honourmsg, OtherMsg.Type, OtherMsg.EndTime)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func AddPost(c *gin.Context) {
	var postmsg pojo.T_article
	if err := c.BindJSON(&postmsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	data, err := logic.NewPostReadService().AddPost(postmsg)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func HonoursImgUpload(c *gin.Context) {

	img, err := c.FormFile("image")
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	honourid, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	url, err := logic.NewFileService().UploadHonourImg(img, honourid)
	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.UploadFail, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, url)
}
