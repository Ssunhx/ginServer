package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	uuid2 "github.com/satori/go.uuid"
)

func getMac() *qbox.Mac {
	mac := qbox.NewMac(QiniuAccessKey, QiniuSecretKey)
	return mac
}

func UploadImg(filestream []byte, filename string) bool {
	mac := getMac()
	// 设置上传到那个空间
	putPolicy := storage.PutPolicy{
		Scope: QiniuBucketName,
	}

	uploadToken := putPolicy.UploadToken(mac)

	// 上传配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = true

	formUploader := storage.NewFormUploader(&cfg)
	// 配置返回的 ret 信息
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": filename,
		},
	}

	// 文件长度
	dataLen := int64(len(filestream))

	// 上传
	err := formUploader.Put(context.Background(), &ret, uploadToken, filename, bytes.NewReader(filestream), dataLen, &putExtra)

	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(ret)
	return true
}

// 生成 uuid
func GenUUID() string {

	u2 := uuid2.NewV4()
	return u2.String()
}

// 删除图片
func DeleteImg(filename string) bool {
	mac := getMac()
	cfg := storage.Config{
		UseHTTPS: true,
		Zone:     &storage.ZoneHuanan,
	}

	bucketmanager := storage.NewBucketManager(mac, &cfg)

	err := bucketmanager.Delete(QiniuBucketName, filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
