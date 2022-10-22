package main

import "flag"

func main() {
	confFile := flag.String("f", "config.yaml", "请指定配置文件")
	flag.Parse()
	conf := InitConfig(*confFile)
	RunCron(conf)
}
