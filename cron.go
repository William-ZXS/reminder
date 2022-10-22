package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func RunCron(conf *Config) {
	c := cron.New()
	var err error
	for _, event := range conf.Events {
		_, err = c.AddFunc(event.Cron, func() {
			conf.Mail.SendMail(event.Spec, "")
		})
	}
	if err != nil {
		fmt.Println("error", err)
	}
	// 启动一个新的 goroutine 做循环检测
	c.Start()
	select {}
}
