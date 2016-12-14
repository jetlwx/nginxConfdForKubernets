package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	LogLevel = beego.AppConfig.String("logLevel")
	Logson   = beego.AppConfig.String("logs")
)

//
//error type to string，方便重定向输出
func CustomerErr(err error) string {
	if err != nil {
		return fmt.Sprintf("%s", err.Error())
	}
	return ""
}

//*******************  write logs  ******************************
//返回日志存储名与路径
func Getlogpath(logtype string) (j string) {
	rotateName := logtype + fmt.Sprintf(".%s.%03d", time.Now().Format("2006-01-02"), 1) + ".log"
	logj := map[string]string{}
	logj["filename"] = beego.AppConfig.String("logDir") + rotateName
	loglog, _ := json.Marshal(&logj)
	//fmt.Println("logj[filename]===", string(loglog))
	return string(loglog)

}

func Writelog(log_level string, Logs ...string) {
	if Logson == "on" {
		var log_type string

		if log_level == "D" {
			log_type = "nginxConfdKube"
		} else {
			log_type = "nginxConfdKube"
		}

		//level 日志级别
		var logshow string
		log := logs.NewLogger(100)

		//异步输出
		log.Async()
		log.SetLogger("file", Getlogpath(log_type))
		for i, _ := range Logs {
			logshow = logshow + " " + Logs[i]
		}
		//fmt.Println("Log_C_memory", Log_C_memory)
		switch LogLevel {
		case "D":
			log.Debug(logshow)
		case "E":
			log.Error(logshow)
		}
	}
}
