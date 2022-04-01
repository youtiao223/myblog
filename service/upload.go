package service

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"myBlog/config"
	"myBlog/utils/errors"
)

// UploadFile 文件上传
func UploadFile(file multipart.File, fileSize int64) (string, int) {
	accessKey := config.QiniuConfig.AccessKey
	secretKey := config.QiniuConfig.SecretKey
	bucket := config.QiniuConfig.Bucket
	url := config.QiniuConfig.Url
	// 上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//1小时有效期
	putPolicy.Expires = 3600
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 构建配置类
	cfg := storage.Config{
		// 空间对应的机房
		Zone: &storage.ZoneHuanan,
		// 是否使用https域名
		UseHTTPS: false,
		// 上传是否使用CDN上传加速
		UseCdnDomains: false,
	}
	// 可选配置
	putExtra := storage.PutExtra{}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errors.ErrorUploadFile
	}

	return url + ret.Key, errors.SUCCESS
}
