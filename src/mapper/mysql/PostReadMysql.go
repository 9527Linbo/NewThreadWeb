package mapper

import (
	"NewThread/src/pojo"
)

type PostReadMysql struct{}

func NewPostReadMysql() *PostReadMysql {
	return &PostReadMysql{}
}

func (*PostReadMysql) ReadShareMysql(id string) (*pojo.Post, error) {
	var post pojo.Post
	err := Db.Raw(`SELECT a.id, a.title, a.content, a.update_time, a.create_time, u.username, iu.url 
	FROM t_article a 
	LEFT JOIN t_user u ON a.user_id = u.id 
	LEFT JOIN t_imageuser iu ON u.id = iu.user_id 
	WHERE a.type = 1 AND a.id = ` + id).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostReadMysql) ReadNewsMysql(id string) (*pojo.Post, error) {
	var post pojo.Post
	err := Db.Raw(`SELECT a.id, a.title, a.content, a.update_time, a.create_time, u.username, iu.url 
	FROM t_article a 
	LEFT JOIN t_user u ON a.user_id = u.id 
	LEFT JOIN t_imageuser iu ON u.id = iu.user_id 
	WHERE a.type = 2 AND a.id = ` + id).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostReadMysql) ReadActivityMysql(id string) (*pojo.Post, error) {
	var post pojo.Post
	err := Db.Raw(`SELECT a.id, a.title, a.content, a.update_time, a.create_time, u.username, iu.url 
	FROM t_article a 
	LEFT JOIN t_user u ON a.user_id = u.id 
	LEFT JOIN t_imageuser iu ON u.id = iu.user_id 
	WHERE a.type = 3 AND a.id = ` + id).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostReadMysql) URLPostByPostIdMysql(id string) (*[]pojo.Url, error) {
	var data []pojo.Url
	err := Db.Raw("SELECT url FROM t_imagearticle where article_id = " + id).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
