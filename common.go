package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{}

func makeRequest(url string) (bodyData string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//if no cookie in config file
	if Config.Cookie != "" {
		req.Header.Set("Cookie", Config.Cookie)
	}

	//send request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	//save response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	//returning response into string
	return string(data), err
}

//MonitorChanges is tha main function for monitoring webpages
func MonitorChanges(ticker time.Duration, urls [][2]string) {

	for range time.Tick(time.Minute * ticker) {
		for i := 0; i < len(urls); i++ {
			resp, err := makeRequest(urls[i][0])
			if err != nil {
				log.Println("MonitorChanges error", err)
			}

			log.Println(urls[i][0])
			urls[i][1] = checkChanges(resp, urls[i][1])
		}
		log.Println("--------Checking webpages completed. Next cycle will start after", ticker, "--------")
	}
}

func checkChanges(newData string, oldData string) string {
	if oldData == "" {
		oldData = string(newData)
		log.Println("Writing fresh data")
	} else if newData == oldData {
		log.Println("There is no change")
	} else {
		log.Println("There is a change")
		oldData = newData
	}

	return oldData
}
