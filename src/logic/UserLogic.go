package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"NewThread/src/utils"
	"mime/multipart"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type UserLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewUserService() *UserLogic {
	return &UserLogic{}
}

func (c *UserLogic) UserIcon(userids []int) ([]pojo.UserIcon, error) {
	data, err := mapper.NewUserMysql().UserIcon(userids)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *UserLogic) UserLogin(usermsg pojo.RecvUserMsg) (string, error) {

	//查数据库是否有该用户
	userpwd, err := mapper.NewUserMysql().UserPwd(usermsg.Account)
	if err != nil {
		return "", err
	}

	//解密
	userpwd, err = utils.RsaDecryptBase64(userpwd)
	if err != nil {
		return "", err
	}

	//密码比较
	if strings.Compare(userpwd, usermsg.Pwd) != 0 {
		return "账号或密码错误", nil
	}

	//存在该用户且密码正确，生成token
	tokenstring, err := utils.GenerateToken(usermsg)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func (c *UserLogic) RegisterUser(usermsg pojo.RecvUserMsg) (string, error) {
	//验证账号是否唯一
	pwd, err := mapper.NewUserMysql().UserPwd(usermsg.Account)
	if err != nil {
		return "", err
	}
	if pwd != "" {
		return "账号已存在", nil
	}
	//密码加密
	usermsg.Pwd, err = utils.RsaEncryptBase64(usermsg.Pwd)
	if err != nil {
		return "", err
	}

	//用户名默认是账号

	//存入数据库的user表
	usermsgMysql := pojo.T_user{Username: usermsg.Account, Account: usermsg.Account, Pwd: usermsg.Pwd, CreatTime: time.Now(), UpdateTime: time.Now()}
	_, err = mapper.NewUserMysql().RegisterUser(usermsgMysql, mapper.Db)
	if err != nil {
		return "", err
	}
	return usermsg.Account, nil
}

func (c *UserLogic) AddTeacher(teachermsg pojo.T_teacher, account string, group string, icon *multipart.FileHeader) (teacher pojo.Teacher, err error) {
	//根据账号查找用户id
	teachermsg.Userid, err = mapper.NewUserMysql().SearhcUserId(account)
	if err != nil {
		return pojo.Teacher{}, err
	}
	teacher.Userid = teachermsg.Userid

	//开启事务
	tx := mapper.Db.Begin()

	_mysql := mapper.NewUserMysql()
	teachermsg.CreatTime = time.Now()
	teachermsg.UpdateTime = time.Now()

	//插入t_teacher表
	err = _mysql.InsertTeacher(teachermsg, tx)
	if err != nil {
		//第一个出错不用回滚
		return pojo.Teacher{}, err
	}
	teacher.Name = teachermsg.Name
	teacher.Position = teachermsg.Position

	//修改负责组id（如果有参数group）
	if group != "" {
		err = _mysql.UpdateTGroupTeacherId(group, teachermsg.Userid, tx)
		if err != nil {
			// 返回err 会自动回滚事务
			tx.Rollback()
			return pojo.Teacher{}, err
		}
		teacher.Group = group
	}

	//上传头像
	url, fileuuid, err := utils.Upload_Simple_File_Clinet_to_OSS(icon, viper.GetString("File.OSSIconPath"))
	if err != nil {
		// 返回err 会自动回滚事务
		tx.Rollback()
		return pojo.Teacher{}, err
	}
	//插入数据库字段
	err = mapper.NewFileMysql().InsertIconMesg(url, fileuuid, teachermsg.Userid, mapper.Db)
	if err != nil {
		// 返回err 会自动回滚事务
		tx.Rollback()
		return pojo.Teacher{}, err
	}
	teacher.URL = url
	//上述三点任何一点报错 直接 数据库事务回滚

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return pojo.Teacher{}, err
	}
	return teacher, nil
}

func (c *UserLogic) AddStudent(studentmsg pojo.T_student, account string, group string, icon *multipart.FileHeader) (student pojo.Student, err error) {
	//根据账号查找用户id
	studentmsg.Userid, err = mapper.NewUserMysql().SearhcUserId(account)
	if err != nil {
		return pojo.Student{}, err
	}
	student.Userid = studentmsg.Userid

	_mysql := mapper.NewUserMysql()
	studentmsg.CreatTime = time.Now()
	studentmsg.UpdateTime = time.Now()

	//用组名查询组id
	groupid, err := _mysql.SearchGroupidBygroupname(group)
	if err != nil {
		return pojo.Student{}, err
	}
	student.Group = group
	student.Class = studentmsg.Class
	studentmsg.GroupId = groupid

	//开启事务
	tx := mapper.Db.Begin()

	//插入t_student表
	err = _mysql.InsertStudent(studentmsg, tx)
	if err != nil {
		//第一个出错不用回滚
		return pojo.Student{}, err
	}
	student.Name = studentmsg.Name
	student.Position = studentmsg.Position

	//上传头像
	url, fileuuid, err := utils.Upload_Simple_File_Clinet_to_OSS(icon, viper.GetString("File.OSSIconPath"))
	if err != nil {
		// 返回err 会自动回滚事务
		tx.Rollback()
		return pojo.Student{}, err
	}
	//插入数据库字段
	err = mapper.NewFileMysql().InsertIconMesg(url, fileuuid, studentmsg.Userid, tx)
	if err != nil {
		// 返回err 会自动回滚事务
		tx.Rollback()
		return pojo.Student{}, err
	}
	student.URL = url
	//上述两点任何一点报错 直接 数据库事务回滚
	//提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return pojo.Student{}, err
	}
	return student, nil
}

func (c *UserLogic) AddGraduate(graduatemsg pojo.T_graduate, account string, icon *multipart.FileHeader) (graduate pojo.StudentHonours, err error) {
	if account != "" {
		//根据账号查找用户id
		graduatemsg.Userid, err = mapper.NewUserMysql().SearhcUserId(account)
		if err != nil {
			return pojo.StudentHonours{}, err
		}
	}
	_mysql := mapper.NewUserMysql()

	url, _, err := utils.Upload_Simple_File_Clinet_to_OSS(icon, viper.GetString("File.OSSIconPath"))
	if err != nil {
		return pojo.StudentHonours{}, err
	}
	graduatemsg.Url = url
	graduatemsg.CreatTime = time.Now()
	graduatemsg.UpdateTime = time.Now()

	//插入t_graduate表
	err = _mysql.InsertGraduate(graduatemsg, mapper.Db)
	if err != nil {
		return pojo.StudentHonours{}, err
	}

	graduate.Name = graduatemsg.Name
	graduate.Description = graduatemsg.Description
	graduate.URL = graduatemsg.Url

	return graduate, nil
}

func (c *UserLogic) UserList() (userlist []pojo.User, err error) {
	userlist, err = mapper.NewUserMysql().UserList()
	if err != nil {
		return nil, err
	}
	return userlist, nil
}

func (c *UserLogic) DelUser(userid string) error {
	icon, err := mapper.NewUserMysql().DelUser(userid, mapper.Db)

	//删除阿里云oss里的头像文件
	if err := utils.DeleteIcon(icon); err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (c *UserLogic) DelTeacher(userid string, delAccount bool) error {
	//开启事务回滚
	tx := mapper.Db.Begin()

	//删除老师信息
	err := mapper.NewUserMysql().DelTeacher(userid, tx)
	if err != nil {
		return err
	}
	if !delAccount {
		tx.Commit()
		return nil
	}
	//删除账号
	icon, err := mapper.NewUserMysql().DelUser(userid, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	//删除阿里云oss里的头像文件
	if err := utils.DeleteIcon(icon); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *UserLogic) DelStudent(userid string, delAccount bool) error {
	//开启事务回滚
	tx := mapper.Db.Begin()

	//删除学生信息
	err := mapper.NewUserMysql().DelStudent(userid, tx)
	if err != nil {
		return err
	}
	if !delAccount {
		tx.Commit()
		return nil
	}
	//删除账号
	icon, err := mapper.NewUserMysql().DelUser(userid, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	//删除阿里云oss里的头像文件
	if err := utils.DeleteIcon(icon); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *UserLogic) DelGraduate(graduateId string) error {
	//删除毕业生信息
	icon, err := mapper.NewUserMysql().DelGraduate(graduateId, mapper.Db)
	if err != nil {
		return err
	}
	//删除阿里云oss里的头像文件
	if err := utils.DeleteIcon(icon); err != nil {
		return err
	}
	return nil
}
