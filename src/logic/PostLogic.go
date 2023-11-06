package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"fmt"
)

type PostLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewPostService() *PostLogic {
	return &PostLogic{}
}

func (c *PostLogic) PostInfo() ([]pojo.Group, error) {
	return mapper.NewPostMysql().PostListMysql()
}

func (c *PostLogic) PostShareInfo() ([]pojo.PostShare, error) {
	//查询所有知识分享动态
	data, err := mapper.NewPostMysql().PostShareMysql()
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	//根据老师的User_id循环查找position表.
	for i := range data {
		position, err := mapper.NewGroupMysql().PositionByUserIdMysql(data[i].Userid)
		if err != nil {
			return nil, err
		}
		mes[i].Position = position
	}
	return data, err
}

func (c *PostLogic) PostNewsInfo() ([]pojo.Teacher, error) {
	// //查询所有老师的姓名、负责小组、用户id
	// mes, err := mapper.NewGroupMysql().GroupTeacherListAndGroupMysql()
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Print(mes)
	// //根据老师的User_id循环查找position表.
	// for i := range mes {
	// 	position, err := mapper.NewGroupMysql().PositionByUserIdMysql(mes[i].Userid)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	mes[i].Position = position
	// }
	// return mes, err
}

func (c *PostLogic) PostActivityInfo() ([]pojo.Teacher, error) {
	// //查询所有老师的姓名、负责小组、用户id
	// mes, err := mapper.NewGroupMysql().GroupTeacherListAndGroupMysql()
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Print(mes)
	// //根据老师的User_id循环查找position表.
	// for i := range mes {
	// 	position, err := mapper.NewGroupMysql().PositionByUserIdMysql(mes[i].Userid)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	mes[i].Position = position
	// }
	// return mes, err
}
