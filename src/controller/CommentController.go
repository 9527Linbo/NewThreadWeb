package controller

import (
	"NewThread/src/logic"
	"NewThread/src/result"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentInfo_topThree(c *gin.Context) {
	articleid, err := strconv.Atoi(c.Query("articleid"))
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	data, err := logic.NewCommentService().CommentInfo_topThree(articleid)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}
