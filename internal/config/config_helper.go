package configure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

//SetConfig 快速解码配置文件并给结构体赋值接口
type SetConfig interface {
	SetConfig(path string)
}

//GetConfigFromFile 读取配置文件返回map类型数据
func GetConfigFromFile(path string, config *map[string]interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open config file error:", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Decode config file error:", err)
		return err
	}

	return err
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

func generateFileObj(fileName string) (*os.File, error) {
	var resultPath string
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(exPath)
	parentDir := filepath.Dir(exDir)
	resultPath = filepath.Join(parentDir, "internal", "configs", fileName)

	file, err := os.Open(resultPath)
	return file, err
}

func generateFilePath(fileName string) string {
	var resultPath string
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(exPath)
	parentDir := filepath.Dir(exDir)
	resultPath = filepath.Join(parentDir, "internal", "configs", fileName)

	return resultPath
}
