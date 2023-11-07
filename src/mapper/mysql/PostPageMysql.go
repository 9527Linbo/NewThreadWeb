package mapper

import (
	"NewThread/src/pojo"
	"strconv"
)

type PostMysql struct{}

func NewPostMysql() *PostMysql {
	return &PostMysql{}
}

func (c *PostMysql) PageShareMysql(current_page, page_size string) (*pojo.PageShare, error) {
	var page pojo.PageShare
	var post []pojo.PostListShare
	c_page, err := strconv.Atoi(current_page)
	if err != nil {
		c_page = 1
	}
	p_size, err := strconv.Atoi(page_size)
	if err != nil {
		p_size = 20
	}
	if err := Db.Raw(`SELECT a.id, a.title, a.content, i.url
	FROM t_article a
	LEFT JOIN (
	  SELECT article_id, MIN(id) AS min_id
	  FROM t_imagearticle
	  GROUP BY article_id
	) i_min ON a.id = i_min.article_id
	LEFT JOIN t_imagearticle i ON i_min.article_id = i.article_id AND i_min.min_id = i.id
	WHERE a.type = 3 LIMIT ` + strconv.Itoa((c_page-1)*p_size) + ", " + strconv.Itoa(p_size)).Scan(&post).Error; err != nil {
		return nil, err
	}
	var share_sum int64
	Db.Table("t_article").Where("type = 3").Count(&share_sum)
	page.TotalPageNum = (int(share_sum) + p_size - 1) / p_size
	page.CurrentPage, _ = strconv.Atoi(current_page)
	page.PostList = &post
	return &page, nil
}

func (c *PostMysql) PageNewsMysql(current_page, page_size string) (*pojo.PageNews, error) {
	var page pojo.PageNews
	var post []pojo.PostListNews
	c_page, err := strconv.Atoi(current_page)
	if err != nil {
		c_page = 1
	}
	p_size, err := strconv.Atoi(page_size)
	if err != nil {
		p_size = 20
	}
	if err := Db.Raw(`SELECT a.id, a.title, a.content 
	FROM t_article a
	WHERE a.type = 2 LIMIT ` + strconv.Itoa((c_page-1)*p_size) + ", " + strconv.Itoa(p_size)).Scan(&post).Error; err != nil {
		return nil, err
	}
	var share_sum int64
	Db.Table("t_article").Where("type = 2").Count(&share_sum)
	page.TotalPageNum = (int(share_sum) + p_size - 1) / p_size
	page.CurrentPage, _ = strconv.Atoi(current_page)
	page.PostList = &post
	return &page, nil
}

func (c *PostMysql) PageActivityMysql(current_page, page_size string) (*pojo.PageActivity, error) {
	var page pojo.PageActivity
	var post []pojo.PostListActivity
	c_page, err := strconv.Atoi(current_page)
	if err != nil {
		c_page = 1
	}
	p_size, err := strconv.Atoi(page_size)
	if err != nil {
		p_size = 20
	}
	if err := Db.Raw(`SELECT a.id, a.title, a.content, i.url
	FROM t_article a
	LEFT JOIN (
	  SELECT article_id, MIN(id) AS min_id
	  FROM t_imagearticle
	  GROUP BY article_id
	) i_min ON a.id = i_min.article_id
	LEFT JOIN t_imagearticle i ON i_min.article_id = i.article_id AND i_min.min_id = i.id
	WHERE a.type = 1 LIMIT ` + strconv.Itoa((c_page-1)*p_size) + ", " + strconv.Itoa(p_size)).Scan(&post).Error; err != nil {
		return nil, err
	}
	var share_sum int64
	Db.Table("t_article").Where("type = 1").Count(&share_sum)
	page.TotalPageNum = (int(share_sum) + p_size - 1) / p_size
	page.CurrentPage, _ = strconv.Atoi(current_page)
	page.PostList = &post
	return &page, nil
}

func (c *PostMysql) URLArticleByArticleIdMysql(article_id int) ([]pojo.PostShareUrl, error) {
	var data []pojo.PostShareUrl
	if err := Db.Raw("SELECT url FROM t_imagearticle where article_id = ?", &article_id).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
