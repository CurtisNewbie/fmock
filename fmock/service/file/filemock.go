package file

import (
	"time"

	"github.com/curtisnewbie/gocommon/util"
	"github.com/curtisnewbie/gocommon/web/dto"
	"github.com/gin-gonic/gin"
)

type Tag struct {
	Id int `json:"id"`

	/** name of tag */
	Name string `json:"name"`

	/** when the record is created */
	CreateTime string `json:"createTime"`

	/** who created this record */
	CreateBy string `json:"createBy"`
}

type ListTagsForFileResp struct {
	Paging  dto.Paging `json:"pagingVo"`
	Payload []Tag      `json:"payload"`
}

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

type DirBrief struct {
	Id   int    `json:"id"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/file-service/open/api/file/list", func(c *gin.Context) {
		l := []FileInfoWebVo{}
		l = append(l, FileInfoWebVo{Id: 1, Uuid: "YBA24ZE-GN0NXT-7ZGY45-MONQLF", Name: "Good Movie.mp4", UploaderName: "Zhuangyongj", UploadTime: dto.TTime(time.Now()), SizeInBytes: 1000000, UserGroup: 0, IsOwner: true, FileType: "FILE"})
		l = append(l, FileInfoWebVo{Id: 2, Uuid: "KDKDFBB-GN0NXT-7ZGY45-MONQLF", Name: "Good Movies", UploaderName: "Zhuangyongj", UploadTime: dto.TTime(time.Now()), SizeInBytes: 1000000, UserGroup: 0, IsOwner: true, FileType: "DIR"})

		util.DispatchOkWData(c, ListFileInfoWebVo{Payload: l, Paging: dto.Paging{Total: 2}})
	})

	router.GET("/file-service/open/api/file/extension/name", func(c *gin.Context) {
		l := []string{}
		l = append(l, "mp4", "txt", "zip", "jpg", "jpeg", "gzip", "txt", "png", "mov")
		util.DispatchOkWData(c, l)
	})

	router.GET("/file-service/open/api/file/tag/list/all", func(c *gin.Context) {
		l := []string{}
		l = append(l, "Movie", "Book", "Coding", "Other")
		util.DispatchOkWData(c, l)
	})

	router.POST("/file-service/open/api/file/tag/list-for-file", func(c *gin.Context) {
		l := []Tag{}
		l = append(l, Tag{Id: 1, Name: "Movie", CreateTime: "2022-09-28", CreateBy: "Zhuangyongj"})
		l = append(l, Tag{Id: 2, Name: "Book", CreateTime: "2022-09-28", CreateBy: "Zhuangyongj"})
		l = append(l, Tag{Id: 3, Name: "Coding", CreateTime: "2022-09-28", CreateBy: "Zhuangyongj"})
		l = append(l, Tag{Id: 4, Name: "Other", CreateTime: "2022-09-28", CreateBy: "Zhuangyongj"})
		util.DispatchOkWData(c, ListTagsForFileResp{Payload: l, Paging: dto.Paging{Total: 1}})
	})

	router.GET("/file-service/open/api/file/dir/list", func(c *gin.Context) {
		l := []DirBrief{}
		l = append(l, DirBrief{Id: 1, Uuid: "YBA24ZE-GN0NXT-7ZGY45-XXXXX", Name: "Old Movies"})
		l = append(l, DirBrief{Id: 2, Uuid: "YBA24ZE-GN0NXT-7ZGY45-VVVVV", Name: "Good Movies"})
		util.DispatchOkWData(c, l)
	})
}
