package upload

import (
	"github.com/xiaoWay/blog/config"
	"mime/multipart"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// 根据配置文件中的配置判断文件上传实例
func NewOSS() OSS {
	switch config.Cfg.Upload.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	default:
		return &Local{}
	}
}
