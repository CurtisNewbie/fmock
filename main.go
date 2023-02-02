package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/gocommon/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func mockHandler(dataFile string, url string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if dataFile == "" {
			server.DispatchErrMsgJson(ctx, "Not supported for mocking")
			return
		}

		data, err := ioutil.ReadFile("fmock/service/mockdata/" + dataFile)
		if err != nil {
			logrus.Errorf("unable to read datafile for '%s', %v", url, err)
			server.DispatchErrMsgJson(ctx, "Not supported for mocking")
		}
		ctx.Data(http.StatusOK, "application/json", data)
	}
}

func main() {

	common.DefaultReadConfig(os.Args)

	for i := 1; true; i++ {
		id := fmt.Sprintf("r%d", i)
		url := common.GetPropStr(fmt.Sprintf("mock.routes.%s.url", id))
		if url == "" {
			break
		}

		dataFile := common.GetPropStr(fmt.Sprintf("mock.routes.%s.data", id))
		method := common.GetPropStr(fmt.Sprintf("mock.routes.%s.method", id))
		method = strings.ToUpper(method)

		if method == "GET" {
			server.PubGet(url, mockHandler(dataFile, url))
		}
		if method == "POST" {
			server.PubPost(url, mockHandler(dataFile, url))
		}
		if method == "DELETE" {
			server.PubDelete(url, mockHandler(dataFile, url))
		}
		if method == "PUT" {
			server.PubPut(url, mockHandler(dataFile, url))
		}
	}

	// configure logging
	server.ConfigureLogging()

	// bootstraping
	server.BootstrapServer()
}
