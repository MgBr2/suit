package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/beevik/ntp"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 工具类!!!
// 写入配置文件
func writeConfig() {
	result, err := json.MarshalIndent(config, "", " ")
	checkErr(err)

	err = ioutil.WriteFile(fileName, result, 644)
	checkErr(err)
}

// 读取配置文件
func readConfig() {
	jsonFile, err := ioutil.ReadFile(fileName)
	checkErr(err)
	err = json.Unmarshal(jsonFile, config)
	checkErr(err)
}

func checkConfig() {
Loop:
	v := reflect.ValueOf(config.Phone)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		if f.String() == "" {
			fmt.Println()
			log.Println("检测到手机信息配置缺失，请检查配置文件，在APP中打开此链接，复制并填写整个响应将会自动匹配喵～")
			log.Println("提示: 可以找机器人客服发送此链接，点开即可喵～")
			fmt.Println("\n\thttps://api.bilibili.com/client_info")
			fmt.Println()

			// 手机配置信息为空，做点什么，顺便把 Buvid 也进来
			formatConfig()
			goto Loop
		}
	}

	// 检测版本
	if config.Bili.TvVersion == "" || config.Bili.AppVersion == "" {
		log.Fatalln("版本信息为空，请重新填写喵！")
	}

	// 检测设备虚拟ID
	if config.Cookies.Buvid == "" {
		log.Println("Buvid 为空，建议重新填写喵！")
		time.Sleep(5 * time.Second)
	}

	// 生成 sessionID
	genSessionID()

	// 格式化 TV 版本
	formatTvVersion()

	// 拼接 User-Agent
	spliceUA()

	// 必要参数为空，需要登录
	if config.Cookies.SESSDATA == "" || config.Cookies.BiliJct == "" || config.AccessKey == "" {
		Login()
	}
}

// 检测手机信息配置
func formatConfig() {
	fmt.Print("请输入: ")
	inputReader := bufio.NewReader(os.Stdin)
	ua, err := inputReader.ReadString('\n')
	checkErr(err)

	ua = strings.Replace(ua, "\\/", "/", -1)

	appBuildVersion := regexp.MustCompile(`build/([\da-zA-z.]+)`)
	appInnerVersion := regexp.MustCompile(`innerVer/([\da-zA-z.]+)`)
	android := regexp.MustCompile(`Android ([\da-zA-z.]+)`)
	build := regexp.MustCompile(`Build/([\da-zA-z.]+)`)
	buvid := regexp.MustCompile(`Buvid/([\da-zA-z.]+)`)
	chromeVersion := regexp.MustCompile(`Chrome/([\da-zA-z.]+)`)
	model := regexp.MustCompile(`model/(\S+)`)
	sdkInt := regexp.MustCompile(`sdkInt/([\da-zA-z.]+)`)

	appBuild := appBuildVersion.FindStringSubmatch(ua)[1]
	appVersion := fmt.Sprintf("%v.%v.%v", string(appBuild[0]), appBuild[1:3], string(appBuild[3]))

	config.Bili.AppVersion = appVersion
	config.Bili.AppBuildVersion = appBuild
	config.Bili.AppInnerVersion = appInnerVersion.FindStringSubmatch(ua)[1]
	config.Phone.AndroidVersion = android.FindStringSubmatch(ua)[1]
	config.Phone.Build = build.FindStringSubmatch(ua)[1]
	config.Cookies.Buvid = buvid.FindStringSubmatch(ua)[1]
	config.Phone.ChromeVersion = chromeVersion.FindStringSubmatch(ua)[1]
	config.Phone.DeviceName = model.FindStringSubmatch(ua)[1]
	config.Phone.AndroidApiLevel = sdkInt.FindStringSubmatch(ua)[1]

	writeConfig()
}

// 格式化 TV 版本
func formatTvVersion() {
	arr := strings.Split(config.Bili.TvVersion, ".")
	for k, v := range arr {
		i, _ := strconv.Atoi(v)
		switch k {
		case 0:
			tvVersion += v
		case 1:
			tvVersion += fmt.Sprintf("%02d", i)
		case 2:
			tvVersion += fmt.Sprintf("%03d", i)
		default:
			log.Fatalln("error!")
		}
	}
}

// 生成 bili_trace_id
func genTraceID() string {
	t := time.Now().Unix()
	x := fmt.Sprintf("%x", t)[0:6]
	randBytes := make([]byte, 26/2)
	_, err := rand.Read(randBytes)
	if err != nil {
		return ""
	}
	r := fmt.Sprintf("%x", randBytes)
	xBiliTraceID := fmt.Sprintf("%v%v:%v%v:0:0", r, x, r[16:26], x)
	return xBiliTraceID
}

// 生成 SessionID
func genSessionID() {
	randBytes := make([]byte, 8/2)
	_, err := rand.Read(randBytes)
	if err != nil {
		return
	}
	sessionID = fmt.Sprintf("%x", randBytes)
}

// 计算签名 (tv端)
func loginSign(params map[string]string) string {
	var query string
	var buffer bytes.Buffer

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		//query += fmt.Sprintf("%v=%v&", k, params[k])
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(params[k])
		buffer.WriteString("&")
	}
	query = strings.TrimRight(buffer.String(), "&")
	//fmt.Println(query)
	sign := strMd5(fmt.Sprintf("%v%v", query, "59b43e04ad6965f34319062b478f83dd"))
	//fmt.Println(sign)
	return sign
}

// 计算签名 (app端)
func appSign(params map[string]string) string {
	var query string
	var buffer bytes.Buffer

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		//query += fmt.Sprintf("%v=%v&", k, params[k])
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(params[k])
		buffer.WriteString("&")
	}
	query = strings.TrimRight(buffer.String(), "&")
	//fmt.Println(query)
	sign := strMd5(fmt.Sprintf("%v%v", query, "560c52ccd288fed045859ed18bffd973"))
	//fmt.Println(sign)
	return sign
}

// 计算 MD5
func strMd5(str string) (retMd5 string) {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 拼接 User-agent
func spliceUA() {
	app := strings.NewReplacer("{android_api_level}", config.Phone.AndroidApiLevel,
		"{android_build}", config.Phone.Build, "{appBuildVer}", config.Bili.AppBuildVersion,
		"{chrome_version}", config.Phone.ChromeVersion, "{device}", config.Buy.Device,
		"{innerVer}", config.Bili.AppInnerVersion, "{osVer}", config.Phone.AndroidVersion,
		"{appVer}", config.Bili.AppVersion, "{phone}", config.Phone.DeviceName,
		"{session_id}", sessionID, "{buvid}", config.Cookies.Buvid)

	appUserAgent = app.Replace("Mozilla/5.0 (Linux; Android {osVer}; {phone} Build/{android_build}; wv) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/{chrome_version} Mobile Safari/537.36 " +
		"os/{device} model/{phone} build/{appBuildVer} osVer/{osVer} sdkInt/{android_api_level} network/2 " +
		"BiliApp/{appBuildVer} mobi_app/{device} channel/bili Buvid/{buvid} sessionID/{session_id} " +
		"innerVer/{innerVer} c_locale/zh-Hans_CN s_locale/zh-Hans_CN disable_rcmd/0 {appVer} os/{device} " +
		"model/{phone} mobi_app/{device} build/{appBuildVer} channel/bili innerVer/{innerVer} " +
		"osVer/{osVer} network/2")

	tv := strings.NewReplacer("{tv_version}", config.Bili.TvVersion, "{tv_build_version}", tvVersion,
		"{device}", config.Phone.DeviceName, "{osVer}", config.Phone.AndroidVersion)

	tvUserAgent = tv.Replace("Mozilla/5.0 BiliTV/{tv_version} os/{device} model/{device} " +
		"mobi_app/android_tv_yst build/{tv_build_version} channel/master innerVer/{tv_build_version} " +
		"osVer/{osVer} network/2")
}

// 格式化 Statics
func formatStatistics() {
	static := &Static{
		AppId:    1,
		Platform: 3,
		Version:  config.Bili.AppVersion,
		Abtest:   "",
	}
	s, err := json.Marshal(static)
	statistics = url.QueryEscape(string(s))
	checkErr(err)
}

// 判断登陆状态
func outPrintAccount() {
	if navs.Code == -101 {
		log.Fatalln("帐号未登录，请检查cookies.")
	}

	bp = navs.Data.Wallet.BcoinBalance
	uname := navs.Data.Uname
	log.Printf("登录成功, 当前帐号: %v, B币余额为: %v.", uname, bp)
}

// 输出装扮信息
func outPrintDetail() {
	itemName = details.Data.Name
	strStartTime = details.Data.Properties.SaleTimeBegin
	if strStartTime == "" {
		log.Fatalln("请输入正确的装扮 id 喵!!!")
	}

	s, err := strconv.ParseInt(strStartTime, 10, 64)
	startTime = s
	checkErr(err)

	intNumLimit, err := strconv.Atoi(details.Data.Properties.SaleBuyNumLimit)
	checkErr(err)

	if intNumLimit < intBuyNum {
		log.Fatalln("设置的购买数量超过上限了喵~")
	}

	if details.Data.CurrentActivity.PriceBpForever == 0 {
		p, _ := strconv.ParseFloat(details.Data.Properties.SaleBpForeverRaw, 64)
		price = p / 100
	} else {
		price = details.Data.CurrentActivity.PriceBpForever / 100
	}

	timeLayout := "2006-01-02 15:04:05"
	log.Printf("装扮名称: %v，开售时间: %v.", details.Data.Name, time.Unix(startTime, 0).Format(timeLayout))

	if time.Now().Unix() >= startTime {
		log.Println("请注意，该装扮可能已经开售了喵～")
	}

	if config.Buy.BpEnough == true {
		if price*float64(intBuyNum) > bp {
			log.Fatalf("您没有足够的钱钱，购买此装扮需要 %.2f B币喵.", price)
		}
	} else {
		if price*float64(intBuyNum) > bp {
			log.Printf("请注意，您没有足够的钱钱，购买此装扮需要 %.2f B币喵!!!!!\n", price)
		}
	}
}

// 输出 Rank
func outPutRank() {
	log.Println("当前装扮列表:")
	fmt.Println("")

	if len(rankInfo.Data.Rank) == 0 {
		fmt.Printf("\t当前列表为空，可能有依号出现!!!\n")
		fmt.Println()
		return
	}

	for _, x := range rankInfo.Data.Rank {
		fmt.Printf("\t编号: %v\t拥有者: %v\n", x.Number, x.Nickname)
	}

	fmt.Println("")
}

/*
	NTP 时间同步
	不论是 Win 还是 Linux, 时间都会跑着跑着就偏掉，Mika 必须给你开个协程来帮你校准喵!
*/
func checkNTP() {
	var notice bool

	for {
		task := time.NewTimer(15 * time.Second)

		ntpTime, err := ntp.Time("ntp.aliyun.com")
		n := time.Now()

		checkErr(err)

		diffTime = n.UnixMilli() - ntpTime.UnixMilli()

		if notice == false {
			log.Printf("当前本地时间差: %v ms.", diffTime)
			log.Println("别担心, Mika 会帮你调整的喵~")
			notice = true
		}

		if diffTime >= 1000 || diffTime <= -1000 {
			log.Println("你的本地时间差太多了喵! Mika 觉得你需要做个 NTP 时间同步喵!")
			log.Fatalln("推荐的 NTP 服务器: ntp.aliyun.com")
		}

		// 接近抢购时间，不要影响程序执行
		if n.Unix() > startTime+30 {
			break
		}

		<-task.C
	}
}

// 获取b站时间
func now() {
	result := &Now{}
	clock := resty.New()
	for {
		resp, err := clock.R().
			SetResult(result).
			EnableTrace().
			SetHeader("user-agent", pcUserAgent).
			Get("http://api.bilibili.com/x/report/click/now")

		checkErr(err)

		if result.Data.Now >= startTime-28 {
			//waitTime = r.Request.TraceInfo().TotalTime.Milliseconds()
			//reviceTime := r.ReceivedAt().UnixMicro()
			//timeLayout := "2006-01-02 15:04:05.000000"
			//t := time.Unix(0, reviceTime*int64(time.Microsecond))
			log.Println("停止获取 b 站时间...")

			ti := resp.Request.TraceInfo()

			fmt.Println()
			fmt.Println("  Received At   :", resp.ReceivedAt())
			fmt.Println()
			fmt.Println("  DNSLookup     :", ti.DNSLookup)
			fmt.Println("  ConnTime      :", ti.ConnTime)
			fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
			fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
			fmt.Println("  ServerTime    :", ti.ServerTime)
			fmt.Println("  ResponseTime  :", ti.ResponseTime)
			fmt.Println("  TotalTime     :", ti.TotalTime)
			fmt.Println("  IsConnReused  :", ti.IsConnReused)
			fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
			fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
			fmt.Println("  RequestAttempt:", ti.RequestAttempt)
			fmt.Println()
			log.Println("Performance Test.")
			break
		}
	}
}

// 获取本地时间
func waitToStart() {
	log.Println("正在等待开售...")
	for {
		task := time.NewTimer(1 * time.Millisecond)
		t := time.Now().Unix()
		fmt.Printf("倒计时: %v.\r", formatSecond(startTime-t))
		if t >= startTime-10 {
			go log.Println("准备冻手!!!")
			break
		}
		<-task.C
	}
}

// 格式化 seconds
func formatSecond(seconds int64) string {
	var d, h, m, s int64
	var msg string

	if seconds > SecondsPerDay {
		d = seconds / SecondsPerDay
		h = seconds % SecondsPerDay / SecondsPerHour
		m = seconds % SecondsPerDay % SecondsPerHour / SecondsPerMinute
		s = seconds % 60
		msg = fmt.Sprintf("%v天%v小时%v分%v秒", d, h, m, s)
	} else if seconds > SecondsPerHour {
		h = seconds / SecondsPerHour
		m = seconds % SecondsPerHour / SecondsPerMinute
		s = seconds % 60
		msg = fmt.Sprintf("%v小时%v分%v秒", h, m, s)
	} else if seconds > SecondsPerMinute {
		m = seconds / SecondsPerMinute
		s = seconds % 60
		msg = fmt.Sprintf("%v分%v秒", m, s)
	} else {
		s = seconds
		msg = fmt.Sprintf("%v秒", s)
	}
	return msg
}
