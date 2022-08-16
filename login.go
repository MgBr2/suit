package main

import (
	"encoding/json"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/skip2/go-qrcode"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Login() {
	var mode int
	log.Println("Cookies 必要参数缺失, 需要进行扫码登录，请选择登陆模式 (输入数字).")
	fmt.Println("\n\t1. 终端中生成二维码.")
	fmt.Println("\t2. 当前目录下生成二维码图片.")
	fmt.Println("\t3. APP 打开 URL 登陆.")
	fmt.Println()

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
	u := url.Values{}

	u.Add("appkey", "4409e2ce8ffd12b8")
	u.Add("local_id", config.Cookies.Buvid)
	u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

	data, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), tvSign(u)))
	checkErr(err)

	req, err := http.NewRequest("POST", "https://passport.bilibili.com/x/passport-tv-login/qrcode/auth_code", strings.NewReader(data))
	checkErr(err)

	req.Header.Set("buvid", config.Cookies.Buvid)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("user-agent", tvUserAgent)
	req.Header.Set("session-id", sessionID)
	req.Header.Set("env", "prod")
	req.Header.Set("app-key", "android_tv_yst")
	req.Header.Set("x-bili-trace-id", genTraceID())
	req.Header.Set("x-bili-aurora-eid", config.Bili.XBiliAuroraEid)
	req.Header.Set("x-bili-aurora-zone", "")

	resp, err := login.Do(req)
	checkErr(err)

	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		checkErr(err)
	}(resp.Body)

	err = json.Unmarshal(body, auth)
	checkErr(err)

	if auth.Code != 0 {
		log.Println(string(body))
	}

	qrCodeUrl = auth.Data.Url
	authCode = auth.Data.AuthCode
}

func getLoginInfo() {
	qrCodePoll := &QrcodePoll{}

	for {
		task := time.NewTimer(3 * time.Second)
		u := url.Values{}

		u.Add("appkey", "4409e2ce8ffd12b8")
		u.Add("auth_code", authCode)
		u.Add("local_id", config.Cookies.Buvid)
		u.Add("ts", strconv.FormatInt(time.Now().Unix(), 10))

		data, err := url.QueryUnescape(fmt.Sprintf("%v&sign=%v", u.Encode(), tvSign(u)))
		checkErr(err)

		req, err := http.NewRequest("POST", "https://passport.bilibili.com/x/passport-tv-login/qrcode/poll", strings.NewReader(data))
		checkErr(err)

		req.Header.Set("buvid", config.Cookies.Buvid)
		req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=utf-8")
		req.Header.Set("user-agent", tvUserAgent)
		req.Header.Set("session-id", sessionID)
		req.Header.Set("env", "prod")
		req.Header.Set("app-key", "android_tv_yst")
		req.Header.Set("x-bili-trace-id", genTraceID())
		req.Header.Set("x-bili-aurora-eid", config.Bili.XBiliAuroraEid)
		req.Header.Set("x-bili-aurora-zone", "")

		resp, err := login.Do(req)
		checkErr(err)

		body, err := io.ReadAll(resp.Body)
		checkErr(err)

		err = json.Unmarshal(body, qrCodePoll)
		checkErr(err)

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

		err = resp.Body.Close()
		checkErr(err)

		<-task.C
	}
}
