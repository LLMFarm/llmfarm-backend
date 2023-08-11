package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/gfile"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func main() {
	infos, _ := cpu.Info()
	modelName := infos[0].ModelName
	// fmt.Println(modelName)

	platform, family, version, _ := host.PlatformInformation()
	// fmt.Println("platform:", platform)
	// fmt.Println("family:", family)
	// fmt.Println("version:", version)
	fingerprintCode := getFingerprintCode(modelName, platform, family, version)
	fmt.Println("fingerprintCode:", fingerprintCode)

	CryptoKey, err := gmd5.EncryptString(fingerprintCode + "llmfarm")
	if err != nil {
		fmt.Println("验证失败:", err)
	}

	binContent := gfile.GetBytes("auth")
	binContent, err = gaes.Decrypt(binContent, []byte(CryptoKey))
	if err != nil {
		fmt.Println("解密失败:", err)
	}

	timeStr := string(binContent)
	tagTimeStr := strings.Split(timeStr, "-")[0]
	endTimeStr := strings.Split(timeStr, "-")[1]

	timeStamp, err := strconv.Atoi(endTimeStr)
	if err != nil {
		fmt.Println("文件转换时间错误:", err)
	}
	tagTimeStamp, err := strconv.Atoi(tagTimeStr)
	if err != nil {
		fmt.Println("文件转换时间错误2:", err)
	}
	nowTimeStamp := time.Now().UnixMilli()

	// 是否改动了服务器时间的判断
	if nowTimeStamp < int64(tagTimeStamp) {
		fmt.Println("验证失败：服务器时间发生过改动:", err)
	}

	if nowTimeStamp > int64(timeStamp) {
		fmt.Println("验证文件过期", err)
	}

	fmt.Println("验证成功：", timeStr)
}

// md5生成机器指纹
func getFingerprintCode(modelName string, platform string, family string, version string) string {
	source := fmt.Sprintf(`%s-%s-%s-%s`, modelName, platform, family, version)
	md5, err := gmd5.EncryptString(source)
	if err != nil {
		fmt.Println("err1", err)
	}
	wudaimaMd5, err := gmd5.EncryptString("llmfarm")
	if err != nil {
		fmt.Println("err2", err)
	}
	fmt.Println("md5: ", md5)
	str, err := gaes.Encrypt([]byte(md5), []byte(wudaimaMd5))
	if err != nil {
		fmt.Println("err3", err)
	}
	newMD5, err := gmd5.EncryptString(string(str))
	if err != nil {
		fmt.Println("err4", err)
	}
	return newMD5
}
