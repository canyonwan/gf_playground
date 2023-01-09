package file

import (
	"context"
	"fmt"
	"gf_playground/internal/consts"
	"gf_playground/internal/dao"
	"gf_playground/internal/model"
	"gf_playground/internal/model/entity"
	"gf_playground/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"time"
)

type sFile struct{}

func init() {
	service.RegisterFile(&sFile{})
}

func (sf *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	if in.File == nil {
		return nil, gerror.New("请上传文件")
	}
	// 定义图上传位置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败")
	}
	// 安全检验: 每人每分钟只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(-time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}

	if count >= consts.FileMaxUploadCountMinute {
		s := fmt.Sprintf("上传次数不能超过%d次", consts.FileMaxUploadCountMinute)
		return nil, gerror.New(s)
	}

	dateDirName := gtime.Now().Format("Ymd")
	// gfile.Join 拼接路径(upload_temp/20220516/xxx.jpg)
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName))
	localFile := fmt.Sprintf("%s/%s/%s", uploadPath, dateDirName, fileName) // todo 应该和gfile.Join()效果一样
	if err != nil {
		return nil, err
	}

	fileKey, err := uploadToQiniu(ctx, localFile, fileName)
	if err != nil {
		return nil, err
	}

	data := entity.FileInfo{
		FileName:     fileName,
		FileUrl:      fileKey,
		FileType:     1,
		FileTypeName: "图片",
		FileSize:     0,
		UserId:       gconv.Int(ctx.Value(consts.CtxAdminId)),
	}

	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.FileUploadOutput{
		Id:           uint(id),
		UserId:       uint(data.UserId),
		FileUrl:      fileKey, // g.Cfg().MustGet(ctx, "qiniu.url").String() 这里不返回了,由前端拼接域名
		FileName:     data.FileName,
		FileSize:     data.FileSize,
		FileType:     int16(data.FileType),
		FileTypeName: data.FileTypeName,
	}, nil
}

// 上传到七牛云
func uploadToQiniu(ctx context.Context, localFile, fileName string) (fileKey string, err error) {
	ak := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	sk := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 // 示例2小时有效期
	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	ret := storage.PutRet{}

	formUploader := storage.NewFormUploader(&cfg)
	err = formUploader.PutFile(context.Background(), &ret, upToken, fileName, localFile, nil)
	if err != nil {
		return "", err
	}
	g.Dump("上传成功: ", ret.Key)
	// 上传成功与失败都要删除本地文件
	err = os.Remove(localFile)
	if err != nil {
		glog.Error(ctx, "删除本地文件失败: %s", err)
	}
	return ret.Key, nil
}
