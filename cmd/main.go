package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
	"wfuzz-ttoc/log"
	"wfuzz-ttoc/payload"
	"wfuzz-ttoc/wfuzz"
)

func start() {
	fmt.Println("            __                  _____ _             \n __      __/ _|_   _ ________  |_   _| |_ ___   ___ \n \\ \\ /\\ / / |_| | | |_  /_  /____| | | __/ _ \\ / __|\n  \\ V  V /|  _| |_| |/ / / /_____| | | || (_) | (__ \n   \\_/\\_/ |_|  \\__,_/___/___|    |_|  \\__\\___/ \\___|\n                                                    ")
	fmt.Println("--Made by Ttoc(https://github.com/try-to-change)\n[This is just a practice project,but if you can improve it, I hope you can tell me.]")
	//启动日志
	err := errors.New("Wfuzz-ttoc is Starting...")
	log.LogError(err)
}

var (
	targetUrl   string
	payloads    string
	concurrency int
	showSuccess string
	outputFile  string
)

func init() {
	flag.StringVar(&targetUrl, "u", "", "Target URL")
	flag.StringVar(&payloads, "p", "", "Payloads")
	flag.IntVar(&concurrency, "c", 10, "Number of concurrency")
	flag.StringVar(&showSuccess, "s", "200/404/303", "Website return code")
	flag.StringVar(&outputFile, "o", "", "Output file")
}

func main() {
	start()

	//命令行解析
	flag.Parse()

	//检测是否有命令行参数
	if targetUrl == "" || payloads == "" {
		fmt.Println("Usage: wfuzz -u <url> -p <payloads> [options]")
		flag.PrintDefaults()
		return
	}
	// 解析 payloads
	payloadList, err1 := payload.ParsePayloads(payloads)
	if err1 != nil {
		log.LogError(err1)
		return
	}

	// 执行 wfuzz
	results, err := wfuzz.Wfuzz(targetUrl, payloadList, concurrency, showSuccess)
	if err != nil {
		log.LogError(err)
		return
	}
	//指定导出文件
	if outputFile != "" {
		file, err := os.Create("../out/" + time.Now().Format("15+04+05") + outputFile + ".txt")
		if err != nil {
			log.LogError(err)
			return
		}
		defer file.Close()

		for _, result := range results {
			_, err := file.WriteString(result + "\n")
			if err != nil {
				log.LogError(err)
				return
			}
		}

		fmt.Printf("Results saved to %s\n", outputFile)
	}
}
