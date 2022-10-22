package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Events []*Event
	Mail   *Mail
}

type Event struct {
	Spec string
	Cron string
}

type Mail struct {
	Host     string
	Port     int
	User     string
	Pwd      string
	Receiver []string
}

//func (c Config) String() string {
//	msg := ""
//	for _, e := range c.Events {
//		msg +=
//	}
//}

func InitConfig(file string) *Config {
	if file == "" {
		file = "./config.yaml"
	}
	viper.SetConfigFile(file)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	//var conf *Config
	conf := &Config{}
	err = viper.Unmarshal(conf)

	if err != nil {
		panic(err)
	}
	fmt.Printf("conf:%+v", *conf)
	return conf
}
