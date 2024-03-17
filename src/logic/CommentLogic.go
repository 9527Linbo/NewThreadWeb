package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
)

type CommentLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewCommentService() *CommentLogic {
	return &CommentLogic{}
}

func (c *CommentLogic) CommentInfo_topThree(postid int) ([]pojo.Comment_topthree, error) {

	commentall, err := mapper.NewCommentMysql().CommentTopThreeMysql(postid)

	if err != nil {
		return nil, err
	}

	//数据处理
	var data []pojo.Comment_topthree

	for i := range commentall {

		if commentall[i].RootCommentId == 0 {

			var temp pojo.Comment_topthree

			//times, _ := time.Parse("2006-01-02 15:04:05", commentall[i].CreatTime)

			//commentall[i].CreatTime = humanize.Time(times)

			temp.Comment = commentall[i]

			data = append(data, temp)
		} else {

			for j := range data {

				if data[j].ID == commentall[i].RootCommentId {

					data[j].SubComment = append(data[j].SubComment, commentall[i])

					break
				}
			}
		}
	}

	return data, nil
}

func (c *CommentLogic) CommentInfo_All(postid int) ([]pojo.Comment, error) {

	data, err := mapper.NewCommentMysql().CommentAllMysql(postid)

	if err != nil {
		return nil, err
	}
	return data, nil
}
