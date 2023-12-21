package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/spf13/viper"
)

/*
简单文件上传（大小不超过5G，对并发上传性能要求不高）： 服务器端---->OSS
*/
func Upload_Simple_File_Server_to_OSS() error {
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile("TestDir/test.txt", "C:/Users/9527/Desktop/新建文本文档.txt")
	if err != nil {
		return err
	}
	return nil
}

/*
简单文件上传（大小不超过5G，对并发上传性能要求不高）： 客户端---->服务器端
*/
func Upload_Simple_File_Clinet_to_Server(header *multipart.FileHeader, path string) error {
	//拿到文件名和存储路径
	filename := header.Filename
	path = viper.GetString("File.ESCPath") + path
	fmt.Print(path)
	//创建一个out流
	out, err := os.Create(path + filename)
	if err != nil {
		return err
	}

	//将内容 写入 context
	context, err := header.Open()
	if err != nil {
		return err
	}

	//将读取的内容写到文件中
	_, err = io.Copy(out, context)
	if err != nil {
		return err
	}

	defer context.Close()
	defer out.Close()
	return nil
}
