package messagemgr

import (
	"reflect"
	"regexp"
	"strings"

	configure "github.com/qqbot_zimmer/zimmer/internal/config"
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
	//cqType := configure.GetCQType()
	//生成器
	generator := Generator{
		message: Message{},
		content: MessageContents{},
	}
	generator.Message(messageBody)

	//监听到的源消息
	sourceMessage := generator.message.Content

	behaviors := configure.GetBehaviors()
	var behavior string
	tag, restMsg := matchRegex(sourceMessage)
	if value, ok := (*behaviors)[tag]; ok {
		behavior = value
	}
	sender := Sender{}
	if behavior == "" {
		behavior = "ReplyCommMsg"
	}
	funcArgs := []reflect.Value{reflect.ValueOf(restMsg), reflect.ValueOf(&generator)}
	RunReflectFunc(behavior, funcArgs)
	sender.SendData(&generator.message)
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

//匹配正则表达式
func matchRegex(str string) (string, string) {
	re := regexp.MustCompile(`^#(?P<content>.*)#(?P<rest>.*)$`)
	match := re.FindStringSubmatch(str)
	return match[1], match[2]
}
