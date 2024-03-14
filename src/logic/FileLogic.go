package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"NewThread/src/utils"
	"mime/multipart"

	"github.com/spf13/viper"
)

type FileLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewFileService() *FileLogic {
	return &FileLogic{}
}

func (c *FileLogic) UploadFile(header *multipart.FileHeader, path string, username string, filename string) error {

	//上传文件
	fileuuid, err := utils.Upload_Simple_File_Clinet_to_Server(header, path)
	if err != nil {
		return err
	}

	//插入数据库字段
	err = mapper.NewFileMysql().InsertFileMesg(filename, fileuuid, username)
	if err != nil {
		return err
	}
	return nil
}

func (c *FileLogic) DownloadFile(path string, fileuuid string, filename string, atOSS bool) (data []byte, err error) {

	if atOSS { //文件在OSS里
		data, err = utils.Download_File_OSS(fileuuid)
	} else { //文件在服务器
		path = viper.GetString("File.ESCPath") + path
		data, err = utils.Download_File_Server(path, fileuuid)
	}

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *FileLogic) FileList(path string) ([]pojo.FileList, error) {

	data_OSS, err := utils.FileList_OSS(path)

	if err != nil {
		return nil, err
	}

	basepath := viper.GetString("File.ESCPath")

	data_ESC, err := utils.FileList_ESC(basepath + path)

	if err != nil {
		return nil, err
	}

	data := append(data_ESC, data_OSS...)

	mysql := mapper.NewFileMysql()
	//做个分页查询？
	for i := range data {

		temp, err := mysql.File(data[i].Fileuuid)

		if err != nil {
			return nil, err
		}

		//data[i].AtOSS = temp.AtOSS
		data[i].Filename = temp.Filename
		data[i].Username = temp.Username
	}
	return data, nil
}
