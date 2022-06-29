package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	pcUserAgent      = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * 60
	SecondsPerDay    = SecondsPerHour * 24
)

var (
	fileName, qrCodeUrl, appUserAgent, tvUserAgent,
	sessionID, tvVersion, authCode, statistics,
	orderId, itemName, strStartTime string
	intBuyNum                                int
	startTime, errorTime, fastTime, diffTime int64
	bp, price                                float64
	rankInfo                                 *Rank
	config                                   = &Config{}
	client                                   = resty.New()
	login                                    = resty.New()
)

// 初始化
func init() {
	var itemID, buyNum string
	var timeBefore int

	// 从命令行读取配置文件
	flag.StringVar(&fileName, "c", "./config.json", "Path to config file.")

	// 从命令行读取装扮id
	flag.StringVar(&itemID, "i", "", "The suit id you want to buy.")

	// 从命令行读取购买数量
	flag.StringVar(&buyNum, "b", "", "The Number of suit you want to buy.")

	// 设置购买时间
	flag.IntVar(&timeBefore, "t", 0, "Set up your time_before.")
	flag.Parse()

	// 读取配置文件
	readConfig()

	if itemID != "" {
		config.Buy.ItemId = itemID
	}

	if buyNum != "" {
		config.Buy.BuyNum = buyNum
	}

	if timeBefore != 0 {
		config.Buy.TimeBefore = timeBefore
	}

	i, err := strconv.Atoi(config.Buy.BuyNum)
	intBuyNum = i
	checkErr(err)

	// 检测一下配置文件
	checkConfig()

	// 设置 baseURL
	client.SetBaseURL("https://api.bilibili.com/x")

	// 设置 Cookies
	cookies := []*http.Cookie{
		{Name: "SESSDATA", Value: config.Cookies.SESSDATA},
		{Name: "bili_jct", Value: config.Cookies.BiliJct},
		{Name: "DedeUserID", Value: config.Cookies.DedeUserID},
		{Name: "DedeUserID__ckMd5", Value: config.Cookies.DedeUserIDCkMd5},
		{Name: "sid", Value: config.Cookies.Sid},
		{Name: "Buvid", Value: config.Cookies.Buvid},
	}
	client.SetCookies(cookies)

	// 转换 Statistics
	formatStatistics()
}

func main() {
	// 初始化log
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	checkErr(err)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	if intBuyNum != 1 {
		log.Printf("当前设置的购买数量为: %v", config.Buy.BuyNum)
	}

	// 检测时间差
	go checkNTP()

	// 登陆验证
	nav()
	outPrintAccount()

	// 未知
	popup()

	// 获取装扮信息
	detail()
	outPrintDetail()

	// 大聪明出现!!
	asset()
	rank()
	stat()
	coupon()

	// 输出编号列表
	outPutRank()

	// 使用本地时间等待开始，在开售前三十秒结束进程
	//waitToStart()

	// 获取b站时间，在开售前二十八秒结束进程
	//now()

	/*
		两个条件:
		1. 请求 b 站时间时的 HTTP 响应时间
		2. 自行设定的 time_before
	*/
	//time.Sleep(time.Duration(27000-waitTime-int64(config.Buy.TimeBefore)) * time.Millisecond)

	// 需要测试，尚不清楚 b 站的 NTP 服务是否相同
	// 使用本地时间等待开始，并监测 NTP 延迟，在开售前十秒结束进程
	waitToStart()

	// 留一秒给模拟进入购买界面
	time.Sleep(time.Duration(9000-int64(config.Buy.TimeBefore)-diffTime) * time.Millisecond)

	// 时间不等人(草了，这里怎么和平常的不一样啊)
	start := time.NewTimer(1000 * time.Millisecond)
	nav()
	detail()
	go asset()
	go stat()
	go state()
	go rank()
	go coupon()

	// 优化?
	// fakeHeader, fakeData := fake()

	// 结束定时器
	<-start.C

	// 创建订单
	//if !fakeCreate(fakeHeader, fakeData) {
	//	create()
	//}
	create()

	// 追踪订单
	tradeQuery()

	// 查询余额
	nav()
	wallet()

	// 查询编号
	suitAsset()

}

func checkErr(err error) {
	if err != err {
		log.Fatalln(err)
	}
}
