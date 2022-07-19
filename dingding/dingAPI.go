package dingding

import (
	"bytes"
	"encoding/json"
	"gongzitiao/myLog"
	"io/ioutil"
	"net/http"
	"strings"
)

type AccessTokenResult struct {
	ErrCode string `json:"errcode"`
	AccessToken string `json:"access_token"`
	ErrMsg string `json:"errmsg"`
	ExpiresIn string `json:"expires_in"`
}

type UserIdResult struct {
	ErrCode string `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	RequestId string `json:"request_id"`
	Result map[string]string `json:"result"`
}

// 文本消息
type MsgData struct {
	MsgType string `json:"msgtype"`
	Text map[string]string `json:"text"`
}

// markdown 消息
type MkMsgData struct {
	MsgType string `json:"msgtype"`
	MarkdownData Markdown `json:"markdown"`
}

type Markdown struct {
	Title string `json:"title"`
	Text string `json:"text"`
}


// 发送 markdown 消息结构
type SendDataMarkdown struct {
	AgentId    string  `json:"agent_id"`
	Msg        MkMsgData `json:"msg"`
	UserIdList string  `json:"userid_list"`
}

// 发送文本消息结构
type SendDataText struct {
	AgentId    string  `json:"agent_id"`
	Msg        MsgData `json:"msg"`
	UserIdList string  `json:"userid_list"`
}

type SendMsgRes struct {
	ErrCode string `json:"errcode"`
	TaskId string `json:"task_id"`
	ErrMsg string `json:"errmsg"`
	RequestId string `json:"request_id"`
}

// 调用钉钉接口，获取 accessToken
func GetAccessToken(appkey, appsecret string) (accessToken string) {
	url_1 := "https://oapi.dingtalk.com/gettoken?appkey="
	url_2 := "&appsecret="
	var urlBuild strings.Builder
	urlBuild.WriteString(url_1)
	urlBuild.WriteString(appkey)
	urlBuild.WriteString(url_2)
	urlBuild.WriteString(appsecret)
	url := urlBuild.String()

	// 调用接口获取accessToken
	resp, err := http.Get(url)
	if err != nil {
		myLog.Logger.Printf("请求accessToken失败：%s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		myLog.Logger.Printf("读取accessToken接口响应失败：%s", err)
		return
	}

	res := AccessTokenResult{}
	_ = json.Unmarshal(body,&res)
	if res.ErrMsg != "ok" {
		myLog.Logger.Printf("请求accessToken响应数据：%s", string(body))
		return
	}
	return res.AccessToken
}

func GetUserId(phoneNum, accessToken string) (userId string)  {
	url_1 := "https://oapi.dingtalk.com/topapi/v2/user/getbymobile?access_token="
	var urlBuild strings.Builder
	urlBuild.WriteString(url_1)
	urlBuild.WriteString(accessToken)
	url := urlBuild.String()

	// 调用接口获取userId
	data := map[string]string{"mobile":phoneNum}
	bytesData, _ := json.Marshal(data)
	resp, err := http.Post(url,"application/json", bytes.NewReader(bytesData))
	if err != nil {
		myLog.Logger.Printf("请求userId失败：%s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		myLog.Logger.Printf("读取userId接口响应失败：%s", err)
		return
	}

	res := UserIdResult{}
	_ = json.Unmarshal(body,&res)

	if res.ErrMsg != "ok" {
		myLog.Logger.Printf("请求 userId 响应数据：%s", string(body))
		return
	}

	return res.Result["userid"]
}



func SendMessage(accessToken, userId, AgentId string, msg MsgData)  {
	url_1 := "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token="
	var urlBuild strings.Builder
	urlBuild.WriteString(url_1)
	urlBuild.WriteString(accessToken)
	url := urlBuild.String()

	sendData := SendDataText{
		AgentId:    AgentId,
		Msg:        msg,
		UserIdList: userId,
	}

	bytesData, _ := json.Marshal(sendData)
	resp, err := http.Post(url,"application/json", bytes.NewReader(bytesData))
	if err != nil {
		myLog.Logger.Printf("请求通知消息失败：%s", string(bytesData))
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		myLog.Logger.Printf("读取消息通知接口响应失败：%s", err)
		return
	}

	res := SendMsgRes{}
	_ = json.Unmarshal(body,&res)
	if res.ErrMsg != "ok" {
		myLog.Logger.Printf("发送消息通知响应数据：%s", string(body))
		return
	}
	myLog.Logger.Printf("发送通知消息成功：%s", string(bytesData))
}

