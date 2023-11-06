package mapper

import (
	"NewThread/src/pojo"
)

type PostMysql struct{}

func NewPostMysql() *PostMysql {
	return &PostMysql{}
}

func (c *PostMysql) PostListMysql() ([]pojo.Group, error) {
	// var m []pojo.Group
	// if err := Db.Raw("select id,name,description from t_group").Scan(&m).Error; err != nil {
	// 	return []pojo.Group{}, err
	// }
	// return m, nil
}

func (c *PostMysql) PostShareMysql() ([]pojo.PostShare, error) {
	var data []pojo.PostShare

	return data, nil
}
