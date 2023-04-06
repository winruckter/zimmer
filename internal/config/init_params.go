package configure

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	onceForParams.Do(initConfigForParams)
	return paramsForRequest
}

//InitConfigForParams 根据配置生成结构体
func initConfigForParams() {
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
	onceForCQType.Do(initConfigForCQType)
	return cqType
}

//InitConfigForCQType 根据配置生成结构体
func initConfigForCQType() {
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

	var p interface{}
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		fmt.Printf("%s (%s) = %v\n", fieldType.Name, fieldType.Type, field.Interface())
	}
}
