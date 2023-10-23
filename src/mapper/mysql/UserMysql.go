package mapper

import "NewThread/src/pojo"

type UserMysql struct{}

func NewUserMysql() *UserMysql {
	return &UserMysql{}
}

func (c *UserMysql) UserSearch() ([]pojo.User, error) {
	var m []pojo.User
	if err := Db.Raw("select * from t_user").Scan(&m).Error; err != nil {
		return []pojo.User{}, err
	}
	return m, nil
}
