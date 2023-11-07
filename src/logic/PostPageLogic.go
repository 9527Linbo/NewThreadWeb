package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PostLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewPostService() *PostLogic {
	return &PostLogic{}
}

// 获取请求中页码和每页文章数量
func GetPageInfo(ctx *gin.Context) (string, string) {
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}
	size := "20"
	return page, size
}

func (c *PostLogic) PageShareInfo(ctx *gin.Context) (*pojo.PageShare, error) {
	page, size := GetPageInfo(ctx)
	//查询所有知识分享文章的标题
	data, err := mapper.NewPostMysql().PageShareMysql(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	return data, err
}

func (c *PostLogic) PageNewsInfo(ctx *gin.Context) (*pojo.PageNews, error) {
	page, size := GetPageInfo(ctx)
	//查询所有新闻文章的标题
	data, err := mapper.NewPostMysql().PageNewsMysql(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	return data, err
}

func (c *PostLogic) PageActivityInfo(ctx *gin.Context) (*pojo.PageActivity, error) {
	page, size := GetPageInfo(ctx)
	//查询所有文化活动文章的标题
	data, err := mapper.NewPostMysql().PageActivityMysql(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	return data, err
}
