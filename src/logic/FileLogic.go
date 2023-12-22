package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/utils"
	"mime/multipart"
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
