package mapper

import (
	"NewThread/src/pojo"
	"errors"

	"gorm.io/gorm"
)

type PostReadMysql struct{}

func NewPostReadMysql() *PostReadMysql {
	return &PostReadMysql{}
}

func (*PostReadMysql) ReadShareMysql(id string) (*pojo.Post, error) {
	var post pojo.Post
	err := Db.Raw(`SELECT a.id, a.title, a.content, a.update_time, a.create_time, u.username, u.url 
	FROM t_article a 
	LEFT JOIN t_user u ON a.user_id = u.id 
	WHERE a.type = 1 AND a.id = ` + id).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostReadMysql) ReadNewsMysql(id string) (*pojo.Post, error) {
	var post pojo.Post
	err := Db.Raw(`SELECT a.id, a.title, a.content, a.update_time, a.create_time, u.username, u.url 
	FROM t_article a 
	LEFT JOIN t_user u ON a.user_id = u.id 
	WHERE a.type = 2 AND a.id = ` + id).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostReadMysql) ReadActivityMysql(id string) (*pojo.Post, error) {
	var post pojo.Post
	err := Db.Raw(`SELECT a.id, a.title, a.content, a.update_time, a.create_time, u.username, u.url 
	FROM t_article a 
	LEFT JOIN t_user u ON a.user_id = u.id 
	WHERE a.type = 3 AND a.id = ` + id).Scan(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostReadMysql) URLPostByPostIdMysql(id string) (*[]pojo.Url, error) {
	var data []pojo.Url
	err := Db.Raw("SELECT url FROM t_articleimage where article_id = " + id).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (*PostReadMysql) InsertArticle(postmsg pojo.T_article, db *gorm.DB) (int, error) {
	m := db.Create(&postmsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return 0, errors.New("Insert---Graduate---Mesg---Error")
	}
	return postmsg.Id, nil

}
