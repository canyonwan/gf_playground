package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type FileUploadInput struct {
	File *ghttp.UploadFile `json:"file" dc:"文件"`
}
type FileUploadOutput struct {
	Id           uint        `json:"id" dc:"文件id"`
	UserId       uint        `json:"user_id" dc:"用户id"`
	FileUrl      string      `json:"fileUrl" dc:"文件url(短地址)"`
	FileName     string      `json:"fileName" dc:"文件名"`
	FileSize     int         `json:"fileSize" dc:"文件大小"`
	FileType     int16       `json:"fileType" dc:"文件类型(枚举)"`
	FileTypeName string      `json:"fileTypeName" dc:"文件类型名称"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
