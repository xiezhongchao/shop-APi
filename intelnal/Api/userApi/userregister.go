package userApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-APi/intelnal/formValidate/userReq"
	"shop-APi/intelnal/service/userSrv"
)

func UserRegister(c *gin.Context) {
	var r userReq.UserRegister
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	register, err := userSrv.UserRegister(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "-1",
			"message": "登录失败",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "登录成功",
		"data":    register,
	})
}
