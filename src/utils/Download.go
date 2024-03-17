package utils

import (
	"NewThread/src/pojo"
	"io"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dustin/go-humanize"
)

// 从OSS里下载文件
func Download_File_OSS(FileName string) ([]byte, error) {
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

func Download_File_Server(path string, fileuuid string) ([]byte, error) {
	file, err := os.Open(path + fileuuid)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(io.Reader(file))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FileList_OSS(path string) ([]pojo.FileList, error) {
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
			temp.Fileuuid = object.Key
			temp.Size = humanize.Bytes(uint64(object.Size))
			temp.UpdateTime = humanize.Time(object.LastModified)
			data = append(data, temp)
		}
	}
	return data, err
}

func FileList_ESC(path string) ([]pojo.FileList, error) {
	var data []pojo.FileList
	file, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range file {
		var temp pojo.FileList
		fileInfo, err := os.Stat(path + f.Name())
		if err != nil {
			return nil, err
		}
		temp.Fileuuid = fileInfo.Name()
		temp.UpdateTime = humanize.Time(fileInfo.ModTime())
		temp.Size = humanize.Bytes(uint64(fileInfo.Size()))
		data = append(data, temp)
	}
	return data, nil
}
