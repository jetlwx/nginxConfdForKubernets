package models

import (
	"crypto/tls"

	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

////get the json data  https or http
func GetjsonData(url string) (HttpStatusCode int, res []byte, er error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r.(error))
		}
	}()
	httpOrhttps := strings.Split(url, ":")
	switch httpOrhttps[0] {
	case "https":
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(url)
		if err != nil {
			er = err
			return 500, []byte(""), er
		}

		Writelog("D", "get url"+url)
		HttpStatusCode = resp.StatusCode
		defer resp.Body.Close()
		if HttpStatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				er = err
				return 500, []byte(""), er
			}
			res = body
		}
		return HttpStatusCode, res, nil

	case "http":
		resp, err := http.Get(url)
		if err != nil {
			return 500, []byte(""), err
		}
		Writelog("D", "get url"+url)

		HttpStatusCode = resp.StatusCode
		defer resp.Body.Close()
		if HttpStatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return 500, []byte(""), err
			}
			res = body
		}
	}
	return HttpStatusCode, res, nil

}

//get the endpoints ip list and port from etcd
func GetendpointsIP(url string) (list []string, e error) {
	code, res, err := GetjsonData(url)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, errors.New("return http code:" + strconv.Itoa(code))
	}
	data := &Endpoints{}
	json.Unmarshal(res, &data)
	//len := len(data.Subsets)

	for _, v := range data.Subsets {
		var iplist []string
		var portlist []string

		for _, v1 := range v.Addresses {
			ip := v1.Ip
			iplist = append(iplist, ip)
		}

		for _, v2 := range v.Ports {
			port := strconv.Itoa(int(v2.Port))
			portlist = append(portlist, port)
		}

		for _, v3 := range iplist {
			for _, v4 := range portlist {
				l := v3 + ":" + v4
				list = append(list, l)
			}
		}
	}
	return list, nil
}
