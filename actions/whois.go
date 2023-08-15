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
	Records [][]Record `json:"records"`
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
	// Iterate through the records to find the "NetName" field
	for _, recordList := range whoisResp.Data.Records {
		for _, record := range recordList {
			if record.Key == "NetName" {
				fmt.Println("NetName:", record.Value)
				return
			}
			if record.Key == "NetRange" {
				fmt.Println("NetRange:", record.Value)
			}
		}
	}

	fmt.Println("NetName not found in records")
	fmt.Println("Response Body:", string(body))
}
