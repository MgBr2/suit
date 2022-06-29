package main

import (
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/skip2/go-qrcode"
	"log"
	"strconv"
	"time"
)

func Login() {
	var mode int
	log.Println("Cookies 必要参数缺失, 需要进行扫码登录，请选择登陆模式 (输入数字).")
	fmt.Println("\n\t1. 终端中生成二维码.")
	fmt.Println("\t2. 当前目录下生成二维码图片.")
	fmt.Println("\t3. APP 打开 URL 登陆.")
	fmt.Println()

	login.SetBaseURL("https://passport.bilibili.com/x/passport-tv-login/qrcode")

Loop:
	_, err := fmt.Scanf("%v", &mode)
	checkErr(err)

	getLoginUrl()

	switch mode {
	case 1:
		obj := qrcodeTerminal.New()
		obj.Get(qrCodeUrl).Print()
		fmt.Println()
	case 2:
		err = qrcode.WriteFile(qrCodeUrl, qrcode.Medium, 256, "./login.png")
		log.Println("已在当前目录下生成二维码，请查看.")
		checkErr(err)
	case 3:
		log.Println("请将此 URL 在 APP 中打开:")
		fmt.Println(qrCodeUrl)
	default:
		fmt.Printf("请重新输入: ")
		goto Loop
	}
	getLoginInfo()
}

func getLoginUrl() {
	auth := &AuthCode{}

	data := map[string]string{
		"appkey":   "4409e2ce8ffd12b8",
		"local_id": config.Cookies.Buvid, // Buvid
		"ts":       strconv.FormatInt(time.Now().Unix(), 10),
	}

	headers := map[string]string{
		"buvid":              config.Cookies.Buvid,
		"user-agent":         tvUserAgent,
		"session-id":         "",
		"env":                "prod",
		"app-key":            "android_tv_yst",
		"x-bili-trace-id":    genTraceID(),
		"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
		"x-bili-aurora-zone": "",
	}

	sign := loginSign(data)
	data["sign"] = sign

	_, err := login.R().
		SetResult(auth).
		SetFormData(data).
		SetHeaders(headers).
		Post("/auth_code")

	checkErr(err)
	qrCodeUrl = auth.Data.Url
	authCode = auth.Data.AuthCode
}

func getLoginInfo() {
	for {
		task := time.NewTimer(3 * time.Second)

		data := map[string]string{
			"appkey":    "4409e2ce8ffd12b8",
			"auth_code": authCode,
			"local_id":  config.Cookies.Buvid, // Buvid
			"ts":        strconv.FormatInt(time.Now().Unix(), 10),
		}

		headers := map[string]string{
			"buvid":              config.Cookies.Buvid,
			"user-agent":         tvUserAgent,
			"session-id":         "",
			"env":                "prod",
			"app-key":            "android_tv_yst",
			"x-bili-trace-id":    genTraceID(),
			"x-bili-aurora-eid":  config.Bili.XBiliAuroraEid,
			"x-bili-aurora-zone": "",
		}

		sign := loginSign(data)
		data["sign"] = sign
		qrCodePoll := &QrcodePoll{}

		_, err := login.R().
			SetResult(qrCodePoll).
			SetFormData(data).
			SetHeaders(headers).
			Post("/poll")

		checkErr(err)
		//fmt.Println(resp)

		if qrCodePoll.Code == 0 {
			config.AccessKey = qrCodePoll.Data.AccessToken
			//fmt.Println("AccessKey: ", qrCodePoll.Data.AccessToken)
			for _, v := range qrCodePoll.Data.CookieInfo.Cookies {
				switch v.Name {
				case "SESSDATA":
					//fmt.Println("SESSDATA: ", v.Value)
					config.Cookies.SESSDATA = v.Value
				case "bili_jct":
					//fmt.Println("bili_jct", v.Value)
					config.Cookies.BiliJct = v.Value
				case "DedeUserID":
					//fmt.Println("DedeUserID", v.Value)
					config.Cookies.DedeUserID = v.Value
				case "DedeUserID__ckMd5":
					//fmt.Println("DedeUserID__ckMd5", v.Value)
					config.Cookies.DedeUserIDCkMd5 = v.Value
				case "sid":
					//fmt.Println("sid: ", v.Value)
					config.Cookies.Sid = v.Value
				}
			}
			writeConfig()
			break
		}

		<-task.C
	}
}
