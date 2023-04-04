package messagemgr

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

//SetOnMessageRouter 设置监听信息路由
func SetOnMessageRouter(r *gin.Engine) {
	r.POST("/", OnMessage)
}

//OnMessage 监听信息
func OnMessage(c *gin.Context) {
	body := c.Request.Body
	messageBody, _ := ioutil.ReadAll(body)
	println(string(messageBody))
	//判断一下是否是file
	isFile := OnMessageForFileMsgs(string(messageBody))
	if !isFile {
		OnMessageForStringMsgs(string(messageBody))
	}
}

//OnMessageForStringMsgs 监听string类型信息(图片类型会监听到stirng类型的CQ码)
func OnMessageForStringMsgs(messageBody string) {
	sourceMessage := gjson.Get(string(messageBody), "message").String()
	if len(sourceMessage) > 1 && sourceMessage[0] == '#' {
		distributer := Distributer{}
		distributer.DitributeMsgToSender(messageBody)
	}
}

//OnMessageForFileMsgs 监听文件类型信息
func OnMessageForFileMsgs(messageBody string) bool {
	exists := gjson.Valid(messageBody) && gjson.Get(messageBody, "file").Exists()
	if exists {
		return true
	}
	return false
}

//CommonProcess 公共快速复读信息
func CommonProcess(rawMsg []byte, context *gin.Context) {
	message := gjson.Get(string(rawMsg), "message").String()

	if len(message) > 1 && message[0] == '#' {
		message = message[1:]
		//uid := gjson.Get(string(rawMsg), "sender.user_id").String()
		nickname := gjson.Get(string(rawMsg), "sender.nickname").String()
		timeInt := gjson.Get(string(rawMsg), "time").Int()
		t := time.Unix(timeInt, 0)

		context.JSON(http.StatusOK, gin.H{
			"reply": "(" + nickname + ")" + "于" + "{" + t.Format("2006-01-02 15:04:05") + "}" + "发送消息：" + message,
		})
	}
}
