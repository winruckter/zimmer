package configure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

//ParamsForRequest 路由及CQTye初始化结构体
type ParamsForRequest struct {
	DefaultPort     string `json:"DefaultPort"`
	CommonURL       string `json:"CommonURL"`
	SendPrivateMsgs string `json:"SendPrivateMsgs"`
	SendGroupMsgs   string `json:"SendGroupMsgs"`
}

var (
	paramsForRequest *ParamsForRequest
	onceForParams    sync.Once
)

//GetParamsForRequest 单例初始化
func GetParamsForRequest() *ParamsForRequest {
	onceForParams.Do(InitConfigForParams)
	return paramsForRequest
}

//InitConfigForParams 根据配置生成结构体
func InitConfigForParams() {
	fileName := "params.json"
	file, err := generateFileObj(fileName)
	if err != nil {
		fmt.Println("Open config file error:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	paramsForRequest = &ParamsForRequest{}
	err = decoder.Decode(paramsForRequest)
	if err != nil {
		fmt.Println("Decode config file error:", err)
		return
	}
}

//CQType 初始化结构体
type CQType struct {
	Face    string `json:"Face"`    //表情
	Record  string `json:"Record"`  //语音
	At      string `json:"At"`      //@
	Forward string `json:"Forward"` //合并消息
	Share   string `json:"Share"`   //分享链接
	Image   string `json:"Image"`   //图片
	Tts     string `json:"Tts"`     //文本转语音
	Poke    string `json:"Poke"`    //戳一戳
}

var (
	cqType        *CQType
	onceForCQType sync.Once
)

//GetCQType 单例初始化
func GetCQType() *CQType {
	onceForCQType.Do(InitConfigForCQType)
	return cqType
}

//InitConfigForCQType 根据配置生成结构体
func InitConfigForCQType() {
	fileName := "cqtype.json"
	file, err := generateFileObj(fileName)
	if err != nil {
		fmt.Println("Open config file error:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cqType = &CQType{}
	err = decoder.Decode(cqType)
	if err != nil {
		fmt.Println("Decode config file error:", err)
		return
	}
}

func generateFileObj(fileName string) (*os.File, error) {
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(exPath)
	parentDir := filepath.Dir(exDir)
	filePath := filepath.Join(parentDir, "internal", "configs", fileName)

	file, err := os.Open(filePath)
	return file, err
}
