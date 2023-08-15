package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type Record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Data struct {
	Records    [][]Record `json:"records"`
	IrrRecords [][]Record `json:"irr_records"`
	Resource   string     `json:"resource"`
	QueryTime  string     `json:"query_time"`
}

type WhoisResponse struct {
	Data Data `json:"data"`
}

func Whois(ip net.IP) {
	url := fmt.Sprintf("https://stat.ripe.net/data/whois/data.json?resource=%s", ip)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var whoisResp WhoisResponse
	err = json.Unmarshal(body, &whoisResp)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Access the extracted data
	for _, irrRecordList := range whoisResp.Data.IrrRecords {
		for _, irrRecord := range irrRecordList {
			switch irrRecord.Key {
			case "route":
				fmt.Println("Route:", irrRecord.Value)
			case "origin":
				fmt.Println("Origin AS:", irrRecord.Value)
			case "descr":
				fmt.Println("Descr:", irrRecord.Value)
			}
		}
	}
	// Access the Resource and QueryTime fields
	fmt.Println("Resource:", whoisResp.Data.Resource)
	fmt.Println("Query Time:", whoisResp.Data.QueryTime)
}
