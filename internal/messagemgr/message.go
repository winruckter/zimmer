package messagemgr

import (
	"fmt"
	"math/rand"
	"time"

	configure "github.com/qqbot_zimmer/zimmer/internal/config"
)

//Message 消息对象
type Message struct {
	Type                 string     //消息类型(private为私人消息，group为群组消息)
	ID                   int64      //消息ID
	From                 int64      //消息来自于谁(目前应该就是本机器人的qq号)
	To                   int64      //消息发送给谁,个人或群qq号
	ToName               string     //消息发送对象的昵称
	Related              int64      //消息的关联对象(群组中，响应某些个人发送的消息，为个人qq号)
	Time                 string     //消息生成的时间
	Content              string     //消息内容
	AdditionalParameters Parameters //消息的附加参数，包括消息是否解析CQ码，及CQ码对应的额外参数
}

//InitMessage 初始化消息
func (m *Message) InitMessage() {
	m.From = 123
}

//GenerateMessage 生成消息
func (m *Message) GenerateMessage(msgType string, to int64, toName string, content string, related ...int64) {
	m.Type = msgType
	m.To = to
	m.ToName = toName
	m.Content = content
	if len(related) > 0 {
		m.Related = related[0]
	}
}

//Parameters 消息的参数(包括CQ码的自定义参数)
type Parameters struct {
	CQ       bool        `default:"false"` //true代表消息为CQ码，默认false
	CQType   string      //CQ码字段
	CQParams interface{} //自定义的CQ码所需的额外参数
}

//SetAdditionalParams 设置附加参数值
func (p *Parameters) SetAdditionalParams(CQ bool, CQType string, CQParams interface{}) {
	p.CQ = CQ
	p.CQType = CQType
	p.CQParams = CQParams
}

//MessageContents 消息内容
type MessageContents struct {
	Value string
}

//GenerateGeneralContent 生成消息内容
func (mc *MessageContents) GenerateGeneralContent(generalMsg string) {
	mc.Value = generalMsg
}

//GenerateCQFaceContent 生成消息内容
func (mc *MessageContents) GenerateCQFaceContent(faceID int64) {
	var num int64
	//999默认不使用传入的id
	if faceID == 999 {
		rand.Seed(time.Now().UnixNano()) // 设置随机数种子
		//随机0-221
		num = rand.Int63n(222)
	} else {
		num = faceID
	}
	cqType := configure.GetCQType()
	CQType := cqType.Face
	//随机生成了一个表情id
	mc.Value = fmt.Sprintf("[CQ:%s,id=%d]", CQType, num) // 使用 Sprintf 函数生成字符串
}

//GenerateCQAtContent 生成消息内容
func (mc *MessageContents) GenerateCQAtContent(to int64, nickname string, isAll bool) {
	cqType := configure.GetCQType()
	CQType := cqType.At

	var name string
	if nickname == "" {
		name = "目标不在群中"
	} else {
		name = nickname
	}
	if isAll {
		//qq号不在群中的时候就会启用name字段
		mc.Value = fmt.Sprintf("[CQ:%s,qq=%s,name=%s]", CQType, "all", name)
	} else {
		mc.Value = fmt.Sprintf("[CQ:%s,qq=%d,name=%s]", CQType, to, name)
	}
}

//GenerateCQShareContent 生成消息内容
func (mc *MessageContents) GenerateCQShareContent(sharedURL string, title string,
	content string /*可选参数*/, image string /*可选参数，图片url*/) {
	cqType := configure.GetCQType()
	CQType := cqType.Share
	mc.Value = fmt.Sprintf("[CQ:%s,url=%s,title=%s,content=%s,image=%s]",
		CQType, sharedURL, title, content, image)
}

//GenerateCQImageContent 生成消息内容
func (mc *MessageContents) GenerateCQImageContent(filePath string /*本地路径或者图片的url*/) {
	cqType := configure.GetCQType()
	CQType := cqType.Image

	/*40000	普通
	40001	幻影
	40002	抖动
	40003	生日
	40004	爱你
	40005	征友*/
	/*图片类型, flash 表示闪照, show 表示秀图, 默认普通图片*/
	if filePath == "" {
		filePath = "file:///C:/Users/lenovo/Desktop/testImage/default.jpg"
	}
	mc.Value = fmt.Sprintf("[CQ:%s,file=%s,type=%s,id=%s]", CQType, filePath, "show", "4000")
}

//GenerateCQTtsContent 生成消息内容
func (mc *MessageContents) GenerateCQTtsContent(text string) {
	//仅支持群聊
	cqType := configure.GetCQType()
	CQType := cqType.Tts

	mc.Value = fmt.Sprintf("[CQ:%s,text=%s]", CQType, text)
}

//GenerateCQPokeContent 生成消息体
func (mc *MessageContents) GenerateCQPokeContent(to int64 /*qq号*/) {
	cqType := configure.GetCQType()
	CQType := cqType.Poke

	mc.Value = fmt.Sprintf("[CQ:%s,qq=%d]", CQType, to)
}
