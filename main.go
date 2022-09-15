package main

import (
	"os"

	"github.com/curtisnewbie/fmock/fmock/service/auth"
	"github.com/curtisnewbie/fmock/fmock/service/file"
	"github.com/curtisnewbie/gocommon/config"
	"github.com/curtisnewbie/gocommon/web/server"
	"github.com/gin-gonic/gin"
)

func main() {
	profile := config.ParseProfile(os.Args)
	conf, e := config.ParseJsonConfig(config.ParseConfigFilePath(os.Args, profile))
	if e != nil {
		panic(e)
	}
	config.SetGlobalConfig(conf)

	// bootstrap web server
	e = server.BootstrapServer(&conf.ServerConf, config.IsProd(profile), func(router *gin.Engine) {
		auth.RegisterRoutes(router)
		file.RegisterRoutes(router)
	})

	if e != nil {
		panic(e)
	}
}
