package utility

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/go-basic/uuid"
)

var LocalHelper = localHelper{}

type localHelper struct {
}

// 上传到OSS
func (s *localHelper) UploadLocal(file multipart.File, filename string, ext string) (string, error) {
	var ctx = gctx.New()
	var saveFileName = uuid.New() + ext
	saveFileName = strings.ReplaceAll(saveFileName, "-", "")
	filePath, _ := g.Cfg().Get(ctx, "localconfig.path")
	saveFilePath := filePath.String() + "/" + "chain/" + saveFileName
	//创建文件夹 chain 如果没有就创建
	_, err := os.Stat(filePath.String() + "/" + "chain/")
	if err != nil {
		err := os.Mkdir(filePath.String()+"/"+"chain/", os.ModePerm)
		if err != nil {
			return "", errors.New("创建文件夹失败")
		}
	}
	//创建文件
	newFile, err := os.Create(saveFilePath)
	if err != nil {
		return "", errors.New("创建文件失败")
	}
	//拷贝文件
	_, err = io.Copy(newFile, file)
	if err != nil {
		return "", errors.New("写入文件失败")
	}
	defer file.Close()
	domain, _ := g.Cfg().Get(ctx, "localconfig.domain")

	res := "http://" + domain.String() + "/" + "chain/" + saveFileName
	return res, nil
}
