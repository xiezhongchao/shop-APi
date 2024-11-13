package global

import (
	"github.com/gin-gonic/gin"
	"shop-APi/intelnal/proto/UserSrv"
)

var (
	Router        *gin.Engine
	UserSrvClient UserSrv.UserClient
)
