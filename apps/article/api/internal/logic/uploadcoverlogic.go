package logic

import (
	"context"
	"fmt"
	"jason-forum/apps/article/api/internal/code"
	"net/http"
	"time"

	"jason-forum/apps/article/api/internal/svc"
	"jason-forum/apps/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10MB
type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover(req *http.Request) (resp *types.UploadCoverResponse, err error) {
	// 定义最大解析文件size并取出文件
	_ = req.ParseMultipartForm(maxFileSize)
	file, handler, err := req.FormFile("cover")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 向阿里云oss发送图片
	// 从环境变量中获取访问凭证。运行本代码示例之前，请先配置环境变量。
	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.BucketName)
	if err != nil {
		logx.Errorf("get bucket failed, err: %v", err)
		return nil, code.GetBucketErr
	}
	objectKey := genFilename(handler.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("put object failed, err: %v", err)
		return nil, code.PutBucketErr
	}

	// 返回图片链接
	return &types.UploadCoverResponse{CoverUrl: genFileURL(objectKey)}, nil
}

func genFilename(filename string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixMilli(), filename)
}

func genFileURL(objectKey string) string {
	return fmt.Sprintf("https://jason_forum.oss-cn-shanghai.aliyuncs.com/%s", objectKey)
}
