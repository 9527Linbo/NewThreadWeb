package utils

import (
	"NewThread/src/pojo"
	"io"
	"mime/multipart"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dustin/go-humanize"
	"github.com/spf13/viper"
)

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

func Upload_Simple_File_Clinet_to_Server(header *multipart.FileHeader) error {
	//拿到文件名和存储路径
	filename := header.Filename
	path := viper.GetString("File.ESCPath")

	//创建一个out流
	out, err := os.Create(path + filename)
	if err != nil {
		return err
	}

	//将内容读入src
	src, err := header.Open()
	if err != nil {
		return err
	}

	//将读取的文件流写到文件中
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	defer src.Close()
	defer out.Close()
	return nil
}

func Download_File(FileName string) ([]byte, error) {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	body, err := bucket.GetObject(FileName)
	if err != nil {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏。
	defer body.Close()

	data, err := io.ReadAll(io.Reader(body))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FileList(path string) ([]pojo.FileList, error) {
	bucket, err := client.Bucket(bucketName)
	var data []pojo.FileList
	if err != nil {
		return nil, err
	}
	// 列举包含指定前缀的文件。默认列举100个文件。
	lsRes, err := bucket.ListObjects(oss.Prefix(path))
	if err != nil {
		return nil, err
	}
	// 打印列举结果。默认情况下，一次返回100条记录。
	for _, object := range lsRes.Objects {
		var temp pojo.FileList
		temp.FileName = object.Key
		temp.Size = humanize.Bytes(uint64(object.Size))
		temp.UpdateTime = humanize.Time(object.LastModified)
		data = append(data, temp)
	}
	return data, err
}
