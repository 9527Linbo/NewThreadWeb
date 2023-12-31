package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"NewThread/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	path := c.PostForm("path")
	username := c.PostForm("username")
	filename := header.Filename
	err = logic.NewFileService().UploadFile(header, path, username, filename)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.UploadFail, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, nil)
}

func DownloadFile(c *gin.Context) {
	path := c.Query("path")
	filename := c.Query("fileName")
	data, err := utils.Download_File(path + filename)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filename) // 用来指定下载下来的文件名
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(data)
}

func FileList(c *gin.Context) {
	path := c.Query("path")
	data, err := utils.FileList(path)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
