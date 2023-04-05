package messagemgr

import (
	configure "github.com/qqbot_zimmer/zimmer/internal/config"
	"github.com/tidwall/gjson"
)

//CQForAt 给参数结构使用的附加结构
type CQForAt struct {
	targetName string
	isAll      bool
}

//CQForTts 给参数结构使用的附加结构
type CQForTts struct {
	text string
}

//Sender 发送者结构
type Sender struct {
	messageType string
	name        string
	qqID        string
}

// SendData 发送数据
func (s Sender) SendData(msg *Message) {
	paramsForRequest := configure.GetParamsForRequest()
	var url string
	// 创建请求体
	request := Request{}
	request.GenerateBody(msg)
	if msg.Type == "group" {
		url = paramsForRequest.CommonURL + paramsForRequest.SendGroupMsgs
	} else if msg.Type == "private" {
		url = paramsForRequest.CommonURL + paramsForRequest.SendPrivateMsgs
	}

	respBody, err := SendRequest(url, request)
	if err != nil {
		panic(err)
	}
	exists := gjson.Valid(string(respBody)) && gjson.Get(string(respBody), "message_id").Exists()
	if exists {
		msg.ID = gjson.Get(string(respBody), "message_id").Int()
	}

	//打印responses
	println(string(respBody))
}
