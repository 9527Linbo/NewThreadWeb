package utils

import (
	"NewThread/src/pojo"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dustin/go-humanize"
)

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

	flag := false // (不包含目录文件)
	// 打印列举结果。默认情况下，一次返回100条记录。
	for _, object := range lsRes.Objects {
		if !flag {
			flag = true
		} else {
			var temp pojo.FileList
			temp.FileName = object.Key
			temp.Size = humanize.Bytes(uint64(object.Size))
			temp.UpdateTime = humanize.Time(object.LastModified)
			data = append(data, temp)
		}
	}
	return data, err
}
