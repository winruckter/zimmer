package configure

import (
	"encoding/json"
	"fmt"
	"os"
)

//SetConfig 快速解码配置文件并给结构体赋值接口
type SetConfig interface {
	SetConfig(path string)
}

//GetConfigFromFile 读取配置文件返回map类型数据
func GetConfigFromFile(path string) (map[string]interface{}, error) {
	config := make(map[string]interface{})

	file, err := os.Open("path")
	if err != nil {
		fmt.Println("Open config file error:", err)
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Decode config file error:", err)
		return config, err
	}

	return config, err
}

//GetConfigObjectFrom 获取某个具体的值
func GetConfigObjectFrom(config map[string]interface{}, objectName string) map[string]interface{} {
	value := make(map[string]interface{})
	if _, ok := config[objectName]; !ok {
		return value
	}
	return value
}

//ConfigKeyIsExist 判断map中是否存在某key值
func ConfigKeyIsExist(config map[string]interface{}, key string) (string, bool) {
	var value string
	if _, ok := config[key]; !ok {
		return value, false
	} else {
		value = config[key].(string)
		return value, false
	}
}
