package file

import (
	"time"

	"github.com/curtisnewbie/gocommon/util"
	"github.com/curtisnewbie/gocommon/web/dto"
	"github.com/gin-gonic/gin"
)

type ListFileInfoWebVo struct {
	Payload []FileInfoWebVo `json:"payload"`
	Paging  dto.Paging      `json:"pagingVo"`
}

type FileInfoWebVo struct {

	/** id */
	Id int `json:"id"`

	/** file's uuid */
	Uuid string `json:"uuid"`

	/** name of the file */
	Name string `json:"name"`

	/** name of the uploader */
	UploaderName string `json:"uploaderName"`

	/** upload time */
	UploadTime dto.TTime `json:"uploadTime"`

	/** size in bytes */
	SizeInBytes int64 `json:"sizeInBytes"`

	/** the group that the file belongs to, 0-public, 1-private */
	UserGroup int `json:"userGroup"`

	/** Whether current user is the owner of this file */
	IsOwner bool `json:"isOwner"`

	/** file type: FILE, DIR */
	FileType string `json:"fileType"`
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/file-service/open/api/file/list", func(c *gin.Context) {
		l := []FileInfoWebVo{}
		f := FileInfoWebVo{Id: 1, Uuid: "YBA24ZE-GN0NXT-7ZGY45-MONQLF", Name: "Good Movie.mp4", UploaderName: "Zhuangyongj", UploadTime: dto.TTime(time.Now()), SizeInBytes: 1000000, UserGroup: 0, IsOwner: true, FileType: "FILE"}
		l = append(l, f)
		util.DispatchOkWData(c, ListFileInfoWebVo{Payload: l, Paging: dto.Paging{Total: 1}})
	})

	router.GET("/file-service/open/api/file/extension/name", func(c *gin.Context) {
		l := []string{}
		l = append(l, "mp4", "txt", "zip", "jpg", "jpeg")
		util.DispatchOkWData(c, l)
	})

	router.GET("/file-service/open/api/file/tag/list/all", func(c *gin.Context) {
		l := []string{}
		l = append(l, "Movie", "Book", "Coding", "Other")
		util.DispatchOkWData(c, l)
	})

}
