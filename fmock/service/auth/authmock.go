package auth

import (
	"github.com/curtisnewbie/gocommon/util"
	"github.com/gin-gonic/gin"
)

type UserWebVo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/auth-service/open/api/user/login", func(c *gin.Context) {
		util.DispatchOkWData(c, "MOCK")
	})

	router.GET("/auth-service/open/api/user/info", func(c *gin.Context) {
		util.DispatchOkWData(c, UserWebVo{Id: 1, Username: "ZhuangyongjMock", Role: "admin"})
	})

}
