package messagemgr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Request 请求体
type Request struct {
	MessageType string `json:"message_type"` //消息类型为群组还是私人
	TargetID    Target `json:"TargetID"`     //消息发送目标(私人消息为(int64)user_id，群组消息为(int64)group_id)
	Message     string `json:"message"`      //消息内容
	AutoEscape  bool   `json:"auto_escape"`  //是否解析CQ码 false为解析
}

//Target 目标
type Target struct {
	UserID  int64 `json:"user_id"`  //发私人消息时的私人qq号
	GroupID int64 `json:"group_id"` //发群组消息时的群组qq号
}

//GenerateBody 生成请求体接口
type GenerateBody interface {
	GenerateBody(message *Message)
}

//GenerateBody 生成请求体
func (req *Request) GenerateBody(message *Message) {
	// 创建请求体
	req.MessageType = message.Type
	req.Message = message.Content
	//这里是否需要解析CQ码
	req.AutoEscape = message.AdditionalParameters.CQ

	//确定是私发消息还是群消息
	if req.MessageType == "private" {
		req.TargetID.UserID = message.To
	} else if req.MessageType == "group" {
		req.TargetID.GroupID = message.To
	}
}

//AddKeyForStruct 给json结构增加K-V
func AddKeyForStruct(body interface{}, key string, value interface{}) []byte {
	//先将body转化为map[string]类型
	bodyMap := make(map[string]interface{})
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(bodyBytes, &bodyMap)

	//增加字段
	bodyMap[key] = value
	newBodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		panic(err)
	}
	return newBodyBytes
}

//SendRequest 根据url和body发送请求，并返回响应数据
func SendRequest(url string, body Request) ([]byte, error) {
	// 构造符合条件的请求体
	var newBody []byte
	if body.MessageType == "private" {
		//增加group_id字段
		newBody = AddKeyForStruct(body, "user_id", body.TargetID.UserID)
	} else if body.MessageType == "group" {
		//增加group_id字段
		newBody = AddKeyForStruct(body, "group_id", body.TargetID.GroupID)
	}

	// 创建http请求
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(newBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	// 发送http请求
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 解码响应体为json
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// 输出响应
	return respBody, nil
}
