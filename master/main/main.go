package main

import (
	"crontab/master"
	"flag"
	"fmt"
	"runtime"
)

var (
	confFile string
)

// 初始化命令行参数
func initArgs() {
	// master -config ./master.json
	flag.StringVar(&confFile, "config", "./master.json", "指定配置文件")
	flag.Parse()
}

// 初始化线程数
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	initArgs()
	initEnv()

	// 初始化配置文件
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 初始化HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	return

ERR:
	fmt.Println(err)
}
