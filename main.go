package main

import (
	"github.com/astaxie/beego"
	"github.com/jetlwx/nginxConfdForKubernets/models"
	"log"
	"strconv"
	"time"
)

func main() {
	log.Println("[D] reading conf from conf file")
	conffile := beego.AppConfig.String("nginxConfFile")
	tmplfile := beego.AppConfig.String("nginxTemplate")
	checkCmd := beego.AppConfig.String("checkCmd")
	reloadCmd := beego.AppConfig.String("reloadCmd")
	apiserver := beego.AppConfig.String("apiServer")
	list := beego.AppConfig.String("servicelist")
	freshSeconds := beego.AppConfig.String("freshSeconds")
	servicelist := beego.AppConfig.String("servicelist")
	delConfFileWhenCheckFaild := beego.AppConfig.String("delConfFileWhenCheckFaild")

	log.Println("[D] starting to resolve the services list")

	urls := models.GetendpointsList(apiserver, list)
	log.Println("[D] Get the endpoints list:", urls)

	for {
		fresh, err := strconv.ParseInt(freshSeconds, 10, 64)
		if err != nil {
			log.Fatal("[E] the freshSeconds must be interge", err)
		}

		DoJob(urls, conffile, tmplfile, checkCmd, reloadCmd, servicelist, delConfFileWhenCheckFaild)
		log.Println("[D] waiting  to next period ......")

		time.Sleep(time.Duration(fresh) * time.Second)
	}

}

func DoJob(urls []string, confile string, tmplfile string, checkCmd string, reloadCmd string, servicelist string, delConfFileWhenCheckFaild string) {

	models.ModifyTemplate(urls, confile, tmplfile, checkCmd, reloadCmd)

}
