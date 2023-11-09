package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostPageLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewPostPageService() *PostPageLogic {
	return &PostPageLogic{}
}

// 获取请求中的页码和每页文章数量
func GetPageInfo(ctx *gin.Context) (string, string) {
	page := ctx.Query("page")
	if _, err := strconv.Atoi(page); err != nil {
		page = "1"
	}
	size := "20"
	return page, size
}

func (*PostPageLogic) PageShareInfo(ctx *gin.Context) (*pojo.PageShare, error) {
	page, size := GetPageInfo(ctx)
	data, err := mapper.NewPostPageMysql().PageShareMysql(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	return data, err
}

func (*PostPageLogic) PageNewsInfo(ctx *gin.Context) (*pojo.PageNews, error) {
	page, size := GetPageInfo(ctx)
	data, err := mapper.NewPostPageMysql().PageNewsMysql(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	return data, err
}

func (*PostPageLogic) PageActivityInfo(ctx *gin.Context) (*pojo.PageActivity, error) {
	page, size := GetPageInfo(ctx)
	data, err := mapper.NewPostPageMysql().PageActivityMysql(page, size)
	if err != nil {
		return nil, err
	}
	fmt.Print(data)
	return data, err
}
