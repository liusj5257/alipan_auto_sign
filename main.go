package main

import (
	// "autoSign/config"
	"autoSign/platform"
	"strings"
	"os"
	"fmt"
)

func main() {
	// pushPlusToken := config.ConfigInstance.PushPlusToken
	// refreshTokens := config.ConfigInstance.RefreshToken
	// bilibiliCookies := config.ConfigInstance.BilibiliCookie
	// jdCookies := config.ConfigInstance.JdCookie
	// kkCookie := config.ConfigInstance.KKCookie

	pushPlusToken := os.Getenv("PUSHPLUS_TOKEN")
	refreshTokens := os.Getenv("REFRESH_TOKEN")
	bilibiliCookies := os.Getenv("BILIBILI_COOKIE")
	jdCookies := os.Getenv("JD_COOKIE")
	kkCookie := os.Getenv("KK_COOKIE")
	fmt.Println("pushPlusToken:", pushPlusToken)
	fmt.Println("refreshTokens:", refreshTokens)
	fmt.Println("bilibiliCookies:", bilibiliCookies)
	fmt.Println("jdCookies:", jdCookies)
	fmt.Println("kkCookie:", kkCookie)
	if refreshTokens != "null" {
		refreshTokenList := strings.Split(refreshTokens, ",")
		aliCloudDisk := &platform.AliCloudDisk{}
		for _, refreshToken := range refreshTokenList {
			aliCloudDisk.Run(pushPlusToken, refreshToken)
		}
	}
	if bilibiliCookies != "null" {
		bilibiliCookieList := strings.Split(bilibiliCookies, ",")
		bilibili := &platform.Bilibili{}
		for _, bilibiliCookie := range bilibiliCookieList {
			bilibili.Run(pushPlusToken, bilibiliCookie)
		}
	}
	if jdCookies != "null" {
		jdCookiesList := strings.Split(jdCookies, ",")
		jd := &platform.JD{}
		for _, value := range jdCookiesList {
			jd.Run(pushPlusToken, value)
		}
	}
	// if kkCookie != "null" {
	// 	kkCookieList := strings.Split(kkCookie, ",")
	// 	kk := &platform.KK{}
	// 	for _, value := range kkCookieList {
	// 		kk.Run(pushPlusToken, value)
	// 	}
	// }
}
