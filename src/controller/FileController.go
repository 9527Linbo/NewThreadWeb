package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	header, err := c.FormFile("file")
	filename := header.Filename
	path := c.PostForm("path")
	username := c.PostForm("username")

	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	err = logic.NewFileService().UploadFile(header, path, username, filename)
	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.UploadFail, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, nil)
}

func DownloadFile(c *gin.Context) {

	path := c.Query("path")
	fileuuid := c.Query("fileuuid")
	filename := c.Query("filename")

	// 把atOSS  string ----> bool
	atOSS, err := strconv.ParseBool(c.Query("atOSS"))
	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	data, err := logic.NewFileService().DownloadFile(path, fileuuid, filename, atOSS)
	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.DownloadFileResp(c, http.StatusOK, filename, data)
}

func FileList(c *gin.Context) {
	path := c.Query("path")

	data, err := logic.NewFileService().FileList(path)
	if err != nil {
		fmt.Print(err)
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
