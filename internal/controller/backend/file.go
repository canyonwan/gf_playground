package backend

import (
	"context"
	"gf_playground/api/v1/backend"
	"gf_playground/internal/model"
	"gf_playground/internal/service"
)

var File = cFile{}

type cFile struct{}

func (cf *cFile) Upload(ctx context.Context, req *backend.UploadFileReq) (res *backend.UploadFileRes, err error) {
	out, err := service.File().Upload(ctx, model.FileUploadInput{
		File: req.File,
	})
	if err != nil {
		return nil, err
	}
	return &backend.UploadFileRes{
		Id:           out.Id,
		UserId:       out.UserId,
		FileUrl:      out.FileUrl,
		FileName:     out.FileName,
		FileSize:     out.FileSize,
		FileType:     out.FileType,
		FileTypeName: out.FileTypeName,
	}, nil
}
