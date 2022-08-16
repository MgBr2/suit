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
	client                                   = &http.Client{}
	login                                    = &http.Client{}
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

	// 转换 Statistics
	formatStatistics()

	// 初始化 HTTP Client
	initialClient()
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
	state()

	// 设置优惠券
	checkCoupon()

	// 输出编号列表
	outPutRank()

	// 等待开始
	waitToStart()

	// 留一秒给模拟进入购买界面
	time.Sleep(time.Duration(9000-int64(config.Buy.TimeBefore)-diffTime) * time.Millisecond)

	// 时间不等人(草了，这里怎么和平常的不一样啊)
	start := time.NewTimer(1000 * time.Millisecond)
	nav()
	detail()
	go asset()
	go stat()
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
