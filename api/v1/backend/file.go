package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/file/upload" method:"post" tags:"文件上传" mime:"multipart/form-data" summary:"文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择文件" dc:"文件"`
}

type UploadFileRes struct {
	Id           uint   `json:"id" dc:"文件id"`
	FileUrl      string `json:"fileUrl" dc:"文件url(短地址)"`
	FileName     string `json:"fileName" dc:"文件名"`
	FileSize     int    `json:"fileSize" dc:"文件大小"`
	FileType     int16  `json:"fileType" dc:"文件类型(枚举)"`
	FileTypeName string `json:"fileTypeName" dc:"文件类型名称"`
	UserId       uint   `json:"userId" dc:"用户id"`
}
