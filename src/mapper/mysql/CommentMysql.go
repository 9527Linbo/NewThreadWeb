package mapper

import (
	"NewThread/src/pojo"
)

type CommentMysql struct{}

func NewCommentMysql() *CommentMysql {
	return &CommentMysql{}
}

func (c *CommentMysql) CommentTopThreeMysql(articleid int) ([]pojo.Comment, error) {
	var m []pojo.Comment

	err := Db.Raw("SELECT tc.id, tc.content, tc.likeCount, tc.creatTime, tc.rootCommentId,t_user.username FROM ("+
		"SELECT  * FROM t_comment WHERE rootCommentId  IS NULL AND articleid = ? AND isDelete = 0 "+
		"UNION "+
		"SELECT * FROM (SELECT * FROM t_comment WHERE articleid = ? AND isDelete = 0) AS c "+
		"WHERE (SELECT count(*) FROM t_comment WHERE rootCommentId=c.rootCommentId AND id "+
		"<= c.id)<=3 AND rootCommentId  = ANY (SELECT id FROM t_comment WHERE rootCommentId IS NULL)  ORDER BY rootCommentId "+
		") AS tc LEFT JOIN t_user ON tc.userid = t_user.id", articleid, articleid).Scan(&m).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}
