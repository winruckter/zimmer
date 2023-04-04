package messagemgr

import (
	"github.com/qqbot_zimmer/zimmer/internal/config/initparams"
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
	NewRequestURL := initparams.GetInstance()
	var url string
	// 创建请求体
	request := Request{}
	request.GenerateBody(msg)
	if msg.Type == "group" {
		url = initparams.CommonURL + NewRequestURL.SendGroupMsgs
	} else if msg.Type == "private" {
		url = initparams.CommonURL + NewRequestURL.SendPrivateMsgs
	}

	// if msg.AdditionalParameters.CQ {
	// 	CQ = false //false代表需要解析CQ,为false是因为后续字段解析CQ的时候是按false来的
	// 	switch msg.AdditionalParameters.CQType {
	// 	case NewRequestURL.CQType.Face:
	// 		msgObj := GenerateFaceMsg{}
	// 		msgObj.GenerateMessage(999)
	// 		msg = msgObj.message
	// 	case NewRequestURL.CQType.At:
	// 		msgObj := GenerateAtMsg{}
	// 		msgObj.GenerateMessage(param.targetID, param.CQForAt.targetName, param.CQForAt.isAll)
	// 		msg = msgObj.message
	// 	case NewRequestURL.CQType.Poke:
	// 		msgObj := GeneratePokeMsg{}
	// 		msgObj.GenerateMessage(param.targetID)
	// 		msg = msgObj.message
	// 	case NewRequestURL.CQType.Tts:
	// 		msgObj := GenerateTtsMsg{}
	// 		msgObj.GenerateMessage(param.CQForTts.text)
	// 		msg = msgObj.message
	// 	}
	// } else {
	// 	msgObj := GenerateStringMsg{}
	// 	msgObj.GenerateMessage(param.message)
	// 	msg = msgObj.message
	// 	CQ = false
	// }

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
