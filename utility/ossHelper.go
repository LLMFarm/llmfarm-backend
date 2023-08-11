package utility

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-basic/uuid"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var OssHelper = ossHelper{}

type ossHelper struct {
}
type Setting struct {
	endpoint     string `json:"endpoint"`
	accessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessKey"`
	bucket       string `json:"bucket"`
	domain       string `json:"domain"`
}

// 上传到OSS
func (s *ossHelper) UploadOss(file []byte, filename string, ext string) (string, error) {
	var saveFileName = uuid.New() + ext
	saveFileName = strings.ReplaceAll(saveFileName, "-", "")
	saveFileName = "chain/" + saveFileName

	var ctx = gctx.New()
	endpoint, err := g.Cfg().Get(ctx, "ossconfig.endpoint")
	accessKey, err := g.Cfg().Get(ctx, "ossconfig.accessKeyId")
	AccessSecret, err := g.Cfg().Get(ctx, "ossconfig.accessKeySecret")
	bucket, err := g.Cfg().Get(ctx, "ossconfig.bucket")
	domain, err := g.Cfg().Get(ctx, "ossconfig.domain")
	setting := &Setting{
		endpoint:     gconv.String(endpoint),
		accessKey:    gconv.String(accessKey),
		AccessSecret: gconv.String(AccessSecret),
		bucket:       gconv.String(bucket),
		domain:       gconv.String(domain),
	}
	err = s.uploadAliYun(file, saveFileName, setting)
	if err != nil {
		return "", err
	}
	res := "https://" + setting.domain + "/" + saveFileName
	return res, err
}

// 上传到阿里云
func (s *ossHelper) uploadAliYun(file []byte, saveFileName string, setting *Setting) error {
	// 创建OSSClient实例。
	client, err := oss.New(setting.endpoint, setting.accessKey, setting.AccessSecret)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(setting.bucket)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	err = bucket.PutObject(saveFileName, bytes.NewReader(file))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
