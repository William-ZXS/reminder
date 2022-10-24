package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func RunCron(conf *Config) {
	c := cron.New()
	var err error
	for _, event := range conf.Events {
		fmt.Println("添加定时任务\n", event.Spec, event.Cron)
		_, err = c.AddFunc(event.Cron, func() {
			err := conf.Mail.SendMail(event.Spec, "")
			if err != nil {
				fmt.Println("send mail err", err)
			}
		})
		if err != nil {
			panic(err)
		}
	}
	
	// 启动一个新的 goroutine 做循环检测
	c.Start()
	select {}
}
