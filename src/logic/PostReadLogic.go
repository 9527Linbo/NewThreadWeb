package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"strconv"

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
