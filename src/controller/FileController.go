package controller

import (
	"NewThread/src/result"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UploadFile(c *gin.Context) {

	header, err := c.FormFile("upload")
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	//拿到文件名和存储路径
	filename := header.Filename
	path := viper.GetString("File.ESCPath")

	//创建一个out流
	out, err := os.Create(path + filename)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	//将内容读入src
	src, err := header.Open()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	//将读取的文件流写到文件中
	_, err = io.Copy(out, src)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}

	defer src.Close()
	defer out.Close()
	result.CommonResp(c, http.StatusOK, result.Success, result.EmptyData)
}
