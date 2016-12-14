package models

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Data1 struct {
	Service string
	Addr    []string
}

var d1 map[string][]string

var lastdata map[string][]string

//get the service list
func GetendpointsList(apiserver string, list string) (l []string) {
	str := strings.Split(list, ",")
	for _, v := range str {
		l2 := apiserver + "/api/v1/namespaces/" + v
		l = append(l, l2)
	}

	return l
}

//Get the service name from url
func GetServiceName(url string) string {
	s := strings.Split(url, "/")
	str := s[len(s)-1]
	return str
}

func GetDataList(urls []string) map[string][]string {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println("[E] program error --> ", err)
		}
	}()

	var d2 = make(map[string][]string)

	for _, url := range urls {
		if url == "" {
			continue
		}

		svcname := GetServiceName(url)
		iplist, e := GetendpointsIP(url)
		if e != nil {
			log.Println("[E] get the endpoints ip list error --> ", e)
		}

		log.Println("[D] get Service name:", svcname)
		log.Println("[D] get the endpoints ips:", iplist)

		if svcname != "" && len(iplist) > 0 {
			d2[svcname] = iplist
		}

	}
	return d2
}

func ModifyTemplate(urls []string, conffile string, tmplfile string, checkCmd string, reloadCmd string) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println("[E] program error --> ", err)
		}
	}()

	datalist := GetDataList(urls)
	eq := CompareData(lastdata, datalist)
	log.Println("[D]eq=", eq)

	if eq == false {
		tmpl, err := template.ParseFiles(tmplfile)

		if err != nil {
			log.Println("[E] template error  --> ", err)
			return
		}
		log.Println("[D] debug info ", datalist)

		var out io.Writer
		var errout error
		out, errout = os.OpenFile(conffile, os.O_RDWR|os.O_CREATE, 0644)
		if errout != nil {
			log.Println("[E] OpenFile error  --> ", err)
			return
		}

		e2 := tmpl.Execute(out, datalist)
		if e2 != nil {
			log.Println("[E] apply the template error  --> ", e2)
			return
		}
		log.Println("[D] the conf file " + conffile + " has been update")

		checkConfFileOK := Execommand(checkCmd)
		if !checkConfFileOK {
			log.Println("[E] check the conf file faild")
			return
		}
		reloadnginxOK := Execommand(reloadCmd)
		if !reloadnginxOK {
			log.Println("[E] reload the nginx  faild")
			return
		}
		log.Println("[D] nginx reload successful!")
		lastdata = datalist
	}
}

func CompareData(lastdata, newdata map[string][]string) bool {
	log.Println("lastdata=", lastdata)
	log.Println("newdata=", newdata)
	for kn, vn := range newdata {
		lenvvn := len(vn)
		log.Println("[D] len=", lenvvn)
		count := 0
		for _, vvn := range vn {
			for kl, vl := range lastdata {
				if kl == kn {
					for _, vvl := range vl {
						if vvn == vvl {
							count += 1
						}
					}
				}
			}
		}

		log.Println("[D] count=", count)
		if lenvvn != count {
			return false
		}

	}
	return true
}

// func CheckConfFile(command string) bool {
// 	cmdArgs := strings.Fields(command)
// 	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
// 	cmdReader, err := cmd.StdoutPipe()
// 	if err != nil {
// 		log.Println(os.Stderr, "Error creating StdoutPipe for Cmd", err)
// 		return false
// 	}

// 	return true
// }

func Execommand(cmdName string) bool {
	cmdArgs := strings.Fields(cmdName)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		log.Println(string(d))
		return false
	}
	log.Println(string(d))
	return true
}
