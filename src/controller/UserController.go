package controller

import (
	"NewThread/src/logic"
	"NewThread/src/pojo"
	"NewThread/src/result"
	"NewThread/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UsersIcon(c *gin.Context) {

	ids, vis := c.GetQueryArray("userids")
	if !vis {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	data, err := logic.NewUserService().UserIcon(utils.StringToInt(ids))
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

// 登录发放token
func Login(c *gin.Context) {
	var usermsg pojo.RecvUserMsg
	//获取用户参数
	if err := c.ShouldBind(&usermsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	data, err := logic.NewUserService().UserLogin(usermsg)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func Register(c *gin.Context) {
	var usermsg pojo.RecvUserMsg
	//获取用户参数
	if err := c.ShouldBind(&usermsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	data, err := logic.NewUserService().RegisterUser(usermsg)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func AddTeacher(c *gin.Context) {
	var teachermsg pojo.T_teacher
	//获取用户参数
	if err := c.ShouldBind(&teachermsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, err.Error())
		return
	}
	account := c.PostForm("account")
	group := c.PostForm("group")
	icon, err := c.FormFile("icon")
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, err.Error())
		return
	}

	data, err := logic.NewUserService().AddTeacher(teachermsg, account, group, icon)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, err.Error())
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func AddStudent(c *gin.Context) {
	var studentmsg pojo.T_student
	//获取用户参数
	if err := c.ShouldBind(&studentmsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	account := c.PostForm("account")
	group := c.PostForm("group")
	icon, err := c.FormFile("icon")
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	data, err := logic.NewUserService().AddStudent(studentmsg, account, group, icon)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func AddGraduate(c *gin.Context) {
	var graduatemsg pojo.T_graduate
	//获取用户参数
	if err := c.ShouldBind(&graduatemsg); err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}
	account := c.PostForm("account")
	icon, err := c.FormFile("icon")
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	data, err := logic.NewUserService().AddGraduate(graduatemsg, account, icon)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func UserList(c *gin.Context) {
	data, err := logic.NewUserService().UserList()
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, data)
}

func DelUser(c *gin.Context) {
	userid := c.Query("userid")
	if userid == "" {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	err := logic.NewUserService().DelUser(userid)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, result.EmptyData)
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, result.EmptyData)
}

func DelTeacher(c *gin.Context) {
	userid := c.Query("userid")
	delAccount, err := strconv.ParseBool(c.Query("delAccount"))
	if userid == "" || err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	err = logic.NewUserService().DelTeacher(userid, delAccount)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, err.Error())
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, result.EmptyData)
}

func DelStudent(c *gin.Context) {
	userid := c.Query("userid")
	delAccount, err := strconv.ParseBool(c.Query("delAccount"))
	if userid == "" || err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	err = logic.NewUserService().DelStudent(userid, delAccount)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, err.Error())
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, result.EmptyData)
}

func DelGraduate(c *gin.Context) {
	graduateId := c.Query("graduateId")
	if graduateId == "" {
		result.CommonResp(c, http.StatusInternalServerError, result.InvalidParam, result.EmptyData)
		return
	}

	err := logic.NewUserService().DelGraduate(graduateId)
	if err != nil {
		result.CommonResp(c, http.StatusInternalServerError, result.ServerBusy, err.Error())
		return
	}
	result.CommonResp(c, http.StatusOK, result.Success, result.EmptyData)
}
