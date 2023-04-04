package initparams

import (
	"sync"
)

//DefaultPort 默认监听端口
const DefaultPort = "8080"

//CommonURL url
const CommonURL = "http://127.0.0.1:5700/"

//NewRequestURL 路由及CQTye初始化结构体
type NewRequestURL struct {
	SendPrivateMsgs string `default:"send_private_msg"`
	SendGroupMsgs   string `default:"send_group_msg"`
	CQType          CQType
}

var (
	instance *NewRequestURL
	// 互斥锁，保证线程安全
	mutex sync.Mutex
)

//CQType 初始化结构体
type CQType struct {
	Face    string `default:"face"`    //表情
	Record  string `default:"record"`  //语音
	At      string `default:"at"`      //@
	Forward string `default:"forward"` //合并消息
	Share   string `default:"share"`   //分享链接
	Image   string `default:"image"`   //图片
	Tts     string `default:"tts"`     //文本转语音
	Poke    string `default:"poke"`    //戳一戳
}

//GetInstance 单例初始化
func GetInstance() *NewRequestURL {
	// 加锁
	mutex.Lock()
	// 如果instance没有被初始化，则初始化它
	if instance == nil {
		instance = &NewRequestURL{
			SendPrivateMsgs: "send_private_msg",
			SendGroupMsgs:   "send_group_msg",
			CQType: CQType{
				Face:    "face",
				Record:  "record",
				At:      "at",
				Forward: "forward",
				Share:   "share",
				Image:   "image",
				Tts:     "tts",
				Poke:    "poke",
			},
		}
	}
	// 解锁
	mutex.Unlock()

	return instance
}
