package mapper

import "NewThread/src/pojo"

type UserMysql struct{}

func NewUserMysql() *UserMysql {
	return &UserMysql{}
}

func (c *UserMysql) UserIcon(userids []int) ([]pojo.UserIcon, error) {
	var m []pojo.UserIcon
	err := Db.Raw("SELECT t_user.id, username,url FROM t_user LEFT JOIN t_imageuser ON t_user.id = t_imageuser.user_id where t_user.id in ?", userids).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
