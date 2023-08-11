package main

import (
	"fmt"

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
	cipher(fingerprintCode)
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

func cipher(fingerprintCode string) {
	CryptoKey, _ := gmd5.EncryptString(fingerprintCode + "llmfarm")
	// binContent, err := ioutil.ReadFile("authSource")
	binContent := []byte("1690905600000-1722614400000")
	binContent, err := gaes.Encrypt(binContent, []byte(CryptoKey))
	if err != nil {
		panic(err)
	}
	if err := gfile.PutBytes("auth", binContent); err != nil {
		panic(err)
	}
}
