package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Webpage Change Monitor")
	Config = ReadConfig()
	log.Println(Config)

	ticker := Config.TimeInterval
	urls := make([][2]string, len(Config.URLs))
	for i := 0; i < len(Config.URLs); i++ {
		urls[i][0] = Config.URLs[i]
	}

	log.Println("Starting timer", ticker)
	MonitorChanges(ticker, urls)
}
