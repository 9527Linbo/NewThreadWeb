package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PostReadLogic struct{}

func NewPostReadService() *PostReadLogic {
	return &PostReadLogic{}
}

func GetPostId(ctx *gin.Context) string {
	id := ctx.Query("id")
	if _, err := strconv.Atoi(id); err != nil {
		id = "1"
	}
	return id
}

func GetImageURL(postid string) *[]pojo.Url {
	url, err := mapper.NewPostReadMysql().URLPostByPostIdMysql(postid)
	if err != nil {
		url = nil
	}
	return url
}

func (c *PostReadLogic) ReadShareInfo(ctx *gin.Context) (*pojo.Post, error) {
	//查询知识分享文章
	id := GetPostId(ctx)
	post, err := mapper.NewPostReadMysql().ReadShareMysql(id)
	if err != nil {
		return nil, err
	}
	post.PostImgUrl = GetImageURL(id)
	return post, err
}

func (c *PostReadLogic) ReadNewsInfo(ctx *gin.Context) (*pojo.Post, error) {
	//查询新闻文章
	id := GetPostId(ctx)
	post, err := mapper.NewPostReadMysql().ReadNewsMysql(id)
	if err != nil {
		return nil, err
	}
	post.PostImgUrl = GetImageURL(id)
	return post, err
}

func (c *PostReadLogic) ReadActivityInfo(ctx *gin.Context) (*pojo.Post, error) {
	//查询文化活动文章
	id := GetPostId(ctx)
	post, err := mapper.NewPostReadMysql().ReadActivityMysql(id)
	if err != nil {
		return nil, err
	}
	post.PostImgUrl = GetImageURL(id)
	return post, err
}

func (c *PostReadLogic) AddPost(postmsg pojo.T_article) (pojo.Post, error) {

	//完善postmsg
	postmsg.CreatTime = time.Now()
	postmsg.UpdateTime = time.Now()

	//插入数据库T_article表
	postid, err := mapper.NewPostReadMysql().InsertArticle(postmsg, mapper.Db)
	if err != nil {
		return pojo.Post{}, err
	}

	//查询用户信息
	userid := []int{postmsg.UserId}
	user, err := mapper.NewUserMysql().UserIcon(userid)
	if err != nil {
		return pojo.Post{}, err
	}

	//完善返回信息
	return pojo.Post{
		Postid:     int64(postid),
		Title:      postmsg.Title,
		Content:    postmsg.Content,
		UpdateTime: postmsg.UpdateTime,
		CreateTime: postmsg.CreatTime,
		UserName:   user[0].Name,
		UserImgUrl: user[0].URL,
	}, nil

}

func (c *PostReadLogic) CreatPostID(typeid int) (postid int, err error) {

	var postmsg pojo.T_article
	postmsg.Type = typeid

	//插入数据库T_article表
	postid, err = mapper.NewPostReadMysql().InsertArticle(postmsg, mapper.Db)
	if err != nil {
		return 0, err
	}
	return postid, nil
}
