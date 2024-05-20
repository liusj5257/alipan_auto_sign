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
	jdCookies := "__jdu=17158421063591771399313; shshshfpa=b60c2cf9-b76e-ddc8-bae4-73bfcd7585a5-1715842106; shshshfpx=b60c2cf9-b76e-ddc8-bae4-73bfcd7585a5-1715842106; areaId=15; ipLoc-djd=15-1158-46341-46352; autoOpenApp_downCloseDate_auto=1715844035916_1800000; unpl=JF8EAJ5nNSttWRhXURhXGUIYH1kGW1xfHEQBZ2YDUgoMSVEBElYbGhB7XlVdWBRKEx9vYBRUWVNIVQ4eBSsSFHtdVV9fD04fC25kNWRVUCVUSBtsGBJdBhBkXV04SicDaWcBXVhbS1UCGQYTEhZNXVZWVQpLFTNuVwVSbWh7VTUaMhoiWyVcGV5bCE8eBmxnBFNfXENUAx0CGRoYSV1Wblw4SA; __jdv=94967808%7Ckong%7Ct_1003649902_%7Cjingfen%7C0b2e2d3a9e5c40ffb39077fe3458e191%7C1715844046514; warehistory="100051772900,"; autoOpenApp_downCloseDate_autoOpenApp_autoPromptly=1715844071491_1; PCSYCityID=CN_330000_330200_0; thor=C0735CEF52A6E37B9AF4F510A149B41693849647F4AEB669C91F8940C398FFD3A80AA65466908E0352F61D66150EC4B84172A16EE4705FFB60D0FEE296A5EA4CEE2F4F1531EDDD4A8181E00A28586735C02E938D4219AD9D3249FC342E2713EC5D1FC9E26CF36A26BE56173F1024A194344369979D536D7F02C1F81CD93FE08E9BDFAC261BFFCD54E9C8765E3421AFA4; pinId=AG8vApQKbxNA9AXplivx1A; pin=liusj5257; unick=%E6%B5%81%E5%85%89%E4%B8%80%E5%88%B9; ceshi3.com=201; _tp=sZNBMw2WXujSiUvi7%2Bog3g%3D%3D; _pst=liusj5257; source=PC; platform=pc; user-key=57e1d10a-b099-44ba-9392-cf96db3f0ec9; TARGET_UNIT=bjcenter; lpkLoginType=3; 3AB9D23F7A4B3CSS=jdd032BZARY4Q7LT3ZH3XLEK35FCHBNYXV4QJNI2AM3TUHL5ROUZSU2DHX7NIFGKEXAYIY7XQ7SKQ4IAZASALMPHENOSXBIAAAAMPSQXVULYAAAAACLJ3DRJ6G5VLEIX; _gia_d=1; token=67335ce2851e0ff8c6f98e78be0d6fed,3,953432; __tk=61595d146d92d608cfd414f670641394,3,953432; jsavif=1; flash=2_czksrmbp9Rn2z8KmWGXROqkhs8E9Lqq260wkiw7LsbhAs6wcH6ex6sre0_hj0BhDhM2Zlchv2AwE09sU63_8axOlFtfb3ArjesHdsztjSmOfG91l2xOAZrgi5GAkRrRB; shshshfpb=BApXcWUcnl-pA-oLpRSkzfAgtGsh2szR_BlfFJilp9xJ1MnuWWIC2; __jda=222648329.17158421063591771399313.1715842106.1716172794.1716178082.4; __jdc=222648329; __jdb=222648329.6.17158421063591771399313|4.1716178082; cn=0; 3AB9D23F7A4B3C9B=2BZARY4Q7LT3ZH3XLEK35FCHBNYXV4QJNI2AM3TUHL5ROUZSU2DHX7NIFGKEXAYIY7XQ7SKQ4IAZASALMPHENOSXBI"
	kkCookie := os.Getenv("KK_COOKIE")
	fmt.Println("pushPlusToken1:", pushPlusToken)
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
