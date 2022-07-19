package main

import (
	"fmt"
	"gongzitiao/dingding"
	"gongzitiao/excel"
	"gongzitiao/myLog"
	"strings"
	"time"
)

var appKey string = "ding"
var appSecret string = "ZhGB"
var agentId string =  "174"

func dingMessageData(row, title []string) (sendData, phoneNum string) {
	var urlBuild strings.Builder
	for s, field := range row {
		if title[s] == "手机号" {
			phoneNum = field
			continue
		}
		if title[s] == "标题" {
			urlBuild.WriteString(field)
			urlBuild.WriteString("\n \n")
			continue
		}
		if title[s] == "结尾" {
			urlBuild.WriteString(field)
			continue
		}
		urlBuild.WriteString(title[s])
		urlBuild.WriteString(": \t")
		urlBuild.WriteString(field)
		urlBuild.WriteString("\n \n")
	}
	sendData = urlBuild.String()
	return sendData, phoneNum
}


func main() {
	// 暂停 5 秒便于使用者查看和窗口信息
	fmt.Println("程序正在运行中。。。。。。")
	time.Sleep(time.Second * 5)

	// 获取excel表格数据和标题栏字段
	title, excelData := excel.GetTitleAndData()
	if title == nil || excelData == nil {
		myLog.Logger.Println("读取excel表格数据失败，程序终止运行")
		return
	}

	// 调用钉钉接口获取accessToken
	accessToken := dingding.GetAccessToken(appKey, appSecret)
	if accessToken == "" {
		myLog.Logger.Println("获取accessToken失败，无法进行后续操作，程序终止")
		return
	}

	// 循环处理excel表格数据
	for _, row := range excelData{
		// 调用 dingMessageData 函数组织发送消息字符串和提取 phoneNum
		sendData, phoneNum := dingMessageData(row, title)

		// 根据 phoneNum 调用钉钉接口获取钉钉用户的 userId
		userId := dingding.GetUserId(phoneNum, accessToken)
		if userId == "" {
			myLog.Logger.Printf("获取 userId 失败，无法发送工作通知，跳过该用户: %s", row)
			continue
		}

		// 组织结构体数据，用于调用钉钉工作通知接口发送的请求体
		msg := dingding.MsgData{
			MsgType: "text",
			Text: map[string]string{"content":sendData},
		}
		// 调用钉钉接口，发送工作通知
		dingding.SendMessage(accessToken, userId, agentId, msg)
	}

	// 暂停 10 秒便于使用者查看和窗口信息
	fmt.Println("所有通知均已发送，10秒后自动退出。")
	time.Sleep(time.Second * 10)
}