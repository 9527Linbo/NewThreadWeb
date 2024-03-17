package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
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
