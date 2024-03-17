package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentInfo_topThree(c *gin.Context) {
	postid, err := strconv.Atoi(c.Query("postid"))
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	data, err := logic.NewCommentService().CommentInfo_topThree(postid)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func CommentInfo_All(c *gin.Context) {
	commentid, err := strconv.Atoi(c.Query("commentid"))
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	data, err := logic.NewCommentService().CommentInfo_All(commentid)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
