package messagemgr

import (
	"strings"

	"github.com/qqbot_zimmer/zimmer/internal/config/initparams"
	"github.com/tidwall/gjson"
)

//Generator 消息生成器
type Generator struct {
	message Message
	content MessageContents
}

//Message 生成一个可发送的消息
func (g *Generator) Message(messageBody string) {
	g.message.InitMessage()
	messageType := gjson.Get(string(messageBody), "message_type").String()
	//获取信息内容
	content := gjson.Get(string(messageBody), "message").String()
	//去掉开头的'#'
	content = content[1:]
	var to int64

	nickname := gjson.Get(string(messageBody), "sender.nickname").String()
	//timeInt := gjson.Get(string(messageBody), "time").Int()
	//t := time.Unix(timeInt, 0)

	if messageType == "group" {
		to = gjson.Get(string(messageBody), "group_id").Int()
		//userInGroupID = gjson.Get(string(messageBody), "sender.user_id").Int()
	} else if messageType == "private" {
		to = gjson.Get(string(messageBody), "sender.user_id").Int()
	}
	//无论是个人还是群组都会有具体的发送消息的某个人的概念，存储该个人的qq号
	related := gjson.Get(string(messageBody), "sender.user_id").Int()
	//这里暂时将监听到的信息内容传入消息
	g.message.GenerateMessage(messageType, to, nickname, content, related)
}

//Distributer 消息分发器
type Distributer struct {
}

//DitributeMsgToSender 向Sender分发消息(之后可能会有多种Sender，留作扩展)
func (g *Distributer) DitributeMsgToSender(messageBody string) {
	//获取一些初始化参数
	initParams := initparams.GetInstance()
	//生成器
	generator := Generator{
		message: Message{},
		content: MessageContents{},
	}
	generator.Message(messageBody)

	//监听到的源消息
	sourceMessage := generator.message.Content
	//消息的附加参数
	addParams := Parameters{}
	//附加参数的自定义CQ参数(自定义匿名结构体)
	var CQParams interface{}
	//待修改
	//考虑这块从配置文件中配置一些关键词关联一些行为方法,将条件语句内的处理整合成行为函数
	substr1 := []string{"复述:"}
	substr2 := []string{"心情"}
	substr3 := []string{"季默"}
	substr4 := []string{"夹你"}
	sender := Sender{}
	if IsContain(sourceMessage, substr1) {
		idx := strings.Index(sourceMessage, substr4[0])
		addParams.CQType = "tts"
		// 取找到的第一个子字符串的后续的字符串
		sourceMessage := sourceMessage[idx+len(substr4[0]):]

		CQParams = struct {
			text string
		}{text: sourceMessage}
		addParams.SetAdditionalParams(true, initParams.CQType.Tts, CQParams)
		//转CQParams类型
		p, ok := CQParams.(struct{ text string })
		if !ok {
			return
		}
		//生成消息内容
		generator.content.GenerateCQTtsContent(p.text)
		generator.message.Content = generator.content.Value
		sender.SendData(&generator.message)
	} else if IsContain(sourceMessage, substr2) {
		CQParams = struct {
			faceID int64
		}{faceID: 999}
		addParams.SetAdditionalParams(true, initParams.CQType.Face, CQParams)
		generator.content.GenerateCQFaceContent(CQParams.(struct{ faceID int64 }).faceID)
		generator.message.Content = generator.content.Value
		sender.SendData(&generator.message)
	} else if IsContain(sourceMessage, substr3) {
		CQParams = struct {
			to       int64
			isAll    bool
			nickname string
		}{to: generator.message.Related, isAll: false,
			nickname: generator.message.ToName}
		addParams.SetAdditionalParams(true, initParams.CQType.At, CQParams)
		//转CQParams类型
		p, ok := CQParams.(struct {
			to       int64
			isAll    bool
			nickname string
		})
		if !ok {
			return
		}
		generator.content.GenerateCQAtContent(p.to, p.nickname, p.isAll)
		generator.message.Content = generator.content.Value
		sender.SendData(&generator.message)
	} else if IsContain(sourceMessage, substr4) {
		CQParams = struct {
			to int64
		}{to: generator.message.Related}
		addParams.SetAdditionalParams(true, initParams.CQType.Poke, CQParams)
		//转CQParams类型
		p, ok := CQParams.(struct{ to int64 })
		if !ok {
			return
		}
		generator.content.GenerateCQPokeContent(p.to)
		generator.message.Content = generator.content.Value
		sender.SendData(&generator.message)
	} else {
		addParams.SetAdditionalParams(false, "", CQParams)
		generator.content.GenerateGeneralContent("test")
		generator.message.Content = generator.content.Value
		sender.SendData(&generator.message)
	}
}

//IsContain 判断消息中是否含有某关键词
func IsContain(msg string, subStr []string) bool {
	for _, strValue := range subStr {
		if strings.Contains(msg, strValue) {
			return true
		}
	}
	return false
}
