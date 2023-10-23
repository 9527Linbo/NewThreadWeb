package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
)

type UserLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewUserService() *UserLogic {
	return &UserLogic{}
}

func (c *UserLogic) UserInfo() ([]pojo.User, error) {
	return mapper.NewUserMysql().UserSearch()
}
