package messagemgr

import (
	"reflect"

	configure "github.com/qqbot_zimmer/zimmer/internal/config"
)

//Reflecter 反射结构体
type Reflecter struct {
}

//ReplyCommMsg 回复一般字符串
func (r Reflecter) ReplyCommMsg(info string, generator *Generator) {
	println("反射成功")
	//消息的附加参数
	addParams := Parameters{}
	//附加参数的自定义CQ参数(自定义匿名结构体)
	var CQParams interface{}
	addParams.SetAdditionalParams(false, "", CQParams)
	generator.content.GenerateGeneralContent("test")
	generator.message.Content = generator.content.Value
}

//Retell 复述
func (r Reflecter) Retell(info string, generator *Generator) {
	println("反射成功")
	//获取一些初始化参数
	cqType := configure.GetCQType()
	//消息的附加参数
	addParams := Parameters{}
	//附加参数的自定义CQ参数(自定义匿名结构体)
	var CQParams interface{}
	addParams.CQType = "tts"
	CQParams = struct {
		text string
	}{text: info}
	addParams.SetAdditionalParams(true, cqType.Tts, CQParams)
	//转CQParams类型
	p, ok := CQParams.(struct{ text string })
	if !ok {
		return
	}

	//生成消息内容
	generator.content.GenerateCQTtsContent(p.text)
	generator.message.Content = generator.content.Value
}

//Emoji 随机一个表情
func (r Reflecter) Emoji(info string, generator *Generator) {
	println("反射成功")
	//获取一些初始化参数
	cqType := configure.GetCQType()
	//消息的附加参数
	addParams := Parameters{}
	//附加参数的自定义CQ参数(自定义匿名结构体)
	var CQParams interface{}
	CQParams = struct {
		faceID int64
	}{faceID: 999}
	addParams.SetAdditionalParams(true, cqType.Face, CQParams)
	generator.content.GenerateCQFaceContent(CQParams.(struct{ faceID int64 }).faceID)
	generator.message.Content = generator.content.Value
}

//AtPerson @某人
func (r Reflecter) AtPerson(info string, generator *Generator) {
	println("反射成功")
	//获取一些初始化参数
	cqType := configure.GetCQType()
	//消息的附加参数
	addParams := Parameters{}
	//附加参数的自定义CQ参数(自定义匿名结构体)
	var CQParams interface{}
	CQParams = struct {
		to       int64
		isAll    bool
		nickname string
	}{to: generator.message.Related, isAll: false,
		nickname: generator.message.ToName}
	addParams.SetAdditionalParams(true, cqType.At, CQParams)
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
}

//PokePerson 戳一戳某人
func (r Reflecter) PokePerson(info string, generator *Generator) {
	println("反射成功")
	//获取一些初始化参数
	cqType := configure.GetCQType()
	//消息的附加参数
	addParams := Parameters{}
	//附加参数的自定义CQ参数(自定义匿名结构体)
	var CQParams interface{}
	CQParams = struct {
		to int64
	}{to: generator.message.Related}
	addParams.SetAdditionalParams(true, cqType.Poke, CQParams)
	//转CQParams类型
	p, ok := CQParams.(struct{ to int64 })
	if !ok {
		return
	}
	generator.content.GenerateCQPokeContent(p.to)
	generator.message.Content = generator.content.Value
}

//RunReflectFunc 运行反射函数
func RunReflectFunc(funcName string, funcArgs []reflect.Value) {
	reflecter := Reflecter{}
	value := reflect.ValueOf(&reflecter)
	f := value.MethodByName(funcName)
	f.Call(funcArgs)
}
