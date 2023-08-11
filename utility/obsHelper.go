package utility

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"
	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"

	"github.com/go-basic/uuid"
)

var ObsHelper = obsHelper{}

type obsHelper struct {
}

// 上传到OSS
func (s *obsHelper) UploadOss(file multipart.File, filename string, ext string) (string, error) {

	var saveFileName = uuid.New() + ext
	saveFileName = strings.ReplaceAll(saveFileName, "-", "")
	saveFileName = "chain/" + saveFileName

	err := s.uploadAliYun(file, saveFileName)
	if err != nil {
		return "", err
	}
	var ctx = gctx.New()

	domain, _ := g.Cfg().Get(ctx, "obsconfig.domain")

	res := "https://" + domain.String() + "/" + saveFileName
	return res, err
}

// 上传到华为云
func (s *obsHelper) uploadAliYun(file multipart.File, saveFileName string) error {

	var ctx = gctx.New()
	endPoint, _ := g.Cfg().Get(ctx, "obsconfig.endpoint")
	ak, _ := g.Cfg().Get(ctx, "obsconfig.accessKeyId")
	sk, _ := g.Cfg().Get(ctx, "obsconfig.accessKeySecret")
	bucket, _ := g.Cfg().Get(ctx, "obsconfig.bucket")

	var obsClient, err = obs.New(ak.String(), sk.String(), endPoint.String())
	if err == nil {
		input := &obs.PutObjectInput{}
		input.Bucket = bucket.String()
		input.Key = saveFileName
		input.Body = file

		output, err := obsClient.PutObject(input)
		if err == nil {
			fmt.Printf("RequestId:%s\n", output.RequestId)
			fmt.Printf("ETag:%s\n", output.ETag)
		} else if obsError, ok := err.(obs.ObsError); ok {
			fmt.Printf("Code:%s\n", obsError.Code)
			fmt.Printf("Message:%s\n", obsError.Message)
		}

		obsClient.Close()
	}
	return nil
}
