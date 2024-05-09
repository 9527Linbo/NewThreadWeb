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

	path = viper.GetString("File.OSSUploadPath") + path

	//上传文件
	url, fileuuid, err := utils.Upload_Simple_File_Clinet_to_OSS(header, path)
	if err != nil {
		return err
	}

	//插入数据库字段
	err = mapper.NewFileMysql().InsertFileMesg(filename, fileuuid, username, url)
	if err != nil {
		return err
	}
	return nil
}

func (c *FileLogic) DownloadFile(path string, fileuuid string, filename string, atOSS bool) (data []byte, err error) {

	path = viper.GetString("File.OSSUploadPath") + path

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

	basepath := viper.GetString("File.OSSUploadPath")
	data, err := utils.FileList_OSS(basepath + path)

	if err != nil {
		return nil, err
	}

	// basepath := viper.GetString("File.ESCPath")

	// data_ESC, err := utils.FileList_ESC(basepath + path)

	// if err != nil {
	// 	return nil, err
	// }

	// data := append(data_ESC, data_OSS...)

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
		data[i].URL = temp.URL
	}
	return data, nil
}

func (c *FileLogic) UploadIcon(header *multipart.FileHeader, userid int, filename string) (string, error) {
	_mysql := mapper.NewFileMysql()

	//查询是否存在旧头像
	oldfileuuid, err := _mysql.SearchFileUUIDById(userid)
	if err != nil {
		return "", err
	}

	if oldfileuuid != "" {
		//删除旧文件
		if err := utils.DeleteIcon(oldfileuuid); err != nil {
			return "", err
		}
	}

	//上传头像（直接上传到OSS）
	url, fileuuid, err := utils.Upload_Simple_File_Clinet_to_OSS(header, viper.GetString("File.OSSIconPath"))
	if err != nil {
		return "", err
	}

	//插入数据库字段
	err = _mysql.InsertIconMesg(url, fileuuid, userid, mapper.Db)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *FileLogic) UploadPostImg(header *multipart.FileHeader, postid int) (string, error) {
	_mysql := mapper.NewFileMysql()

	//上传图片（直接上传到OSS）
	url, _, err := utils.Upload_Simple_File_Clinet_to_OSS(header, viper.GetString("File.OSSSPostImg"))
	if err != nil {
		return "", err
	}

	//插入数据库字段
	err = _mysql.InsertPostImgMesg(url, postid, mapper.Db)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *FileLogic) UploadHonourImg(header *multipart.FileHeader, honourid int) (string, error) {
	_mysql := mapper.NewFileMysql()

	//上传图片（直接上传到OSS）
	url, _, err := utils.Upload_Simple_File_Clinet_to_OSS(header, viper.GetString("File.OSSHonourImg"))
	if err != nil {
		return "", err
	}

	//插入数据库字段
	err = _mysql.InsertHonourImgMesg(url, honourid, mapper.Db)
	if err != nil {
		return "", err
	}
	return url, nil
}
