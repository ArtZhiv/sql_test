package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	urlRIPE = "http://rest.db.ripe.net/search?query-string="
)

type RipeSearch struct {
	XMLName xml.Name `xml:"whois-resources"`
	// Objects []Object `xml:"service"`
	Objects []Object `xml:"objects"`
}

type Object struct {
	Attributes []string `xml:"attributes"`
}

// RipeRequest ...
// запрос информации по IP
// 185.60.216.52
func RipeRequest() {
	ipAddress := "185.60.216.52"
	url := fmt.Sprintf("%v%v", urlRIPE, ipAddress)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/xml")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(string(body))

	var r RipeSearch

	err = xml.Unmarshal(body, &r)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(r)
	fmt.Println(r.Objects)
}
