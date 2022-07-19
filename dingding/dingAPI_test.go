package dingding

import (
	"fmt"
	"testing"
)

func TestGetAccessTokenT(t *testing.T)  {
	appkey := "ding2jrjjmezvxbfxbr4"
	appsecret := "x2XAzl-wz30Mqwrpbrg55TPPma4kw7cDgXJEh9snAOgt_fKf8ai2tgZKOSD5Nx6L"
	accessToken := GetAccessToken(appkey, appsecret)
	fmt.Printf("this is accessToken: %s", accessToken)
}

func TestGetUserId(t *testing.T) {
	appkey := "ding2jrjjmezvxbfxbr4"
	appsecret := "x2XAzl-wz30Mqwrpbrg55TPPma4kw7cDgXJEh9snAOgt_fKf8ai2tgZKOSD5Nx6L"
	accessToken := GetAccessToken(appkey, appsecret)
	userId := GetUserId("18810602951", accessToken)
	fmt.Printf("this is userId: %s", userId)
}

func TestSendMessage(t *testing.T) {
	appkey := "dingyzgg0cdpokf70gfx"
	appsecret := "ZhGbM4GaSViJiMWh3UxNztoccljAUhmg0uCW66qWeYoLolOxqR97LKPBRRBpSGMB"
	accessToken := GetAccessToken(appkey, appsecret)
	userId := GetUserId("18810602951", accessToken)

	msg := MsgData{
		MsgType: "text",
		Text: map[string]string{"content": "1234556666"},
	}

	AgentId := "1743258061"

	SendMessage(accessToken, userId, AgentId, msg)
}