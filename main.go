package main

import (
	"github.com/curtisnewbie/fmock/fmock/service/auth"
	"github.com/curtisnewbie/fmock/fmock/service/file"
	"github.com/curtisnewbie/gocommon/config"
	"github.com/curtisnewbie/gocommon/web/server"
	"github.com/gin-gonic/gin"
)

func main() {
	profile, conf := config.DefaultParseProfConf()

	// bootstrap web server
	e := server.BootstrapServer(&conf.ServerConf, config.IsProd(profile), func(router *gin.Engine) {
		auth.RegisterRoutes(router)
		file.RegisterRoutes(router)
	})

	if e != nil {
		panic(e)
	}
}
