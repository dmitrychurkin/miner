package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	host := flag.String("host", "https://faucet.ropsten.be/donate", "Defines a remote host of maining service")
	walletAddress := flag.String("addr", "0xeA762878a8Dd131Ecfc46c6A45ED9F1EE9d095B3", "MM wallet address")
	duration := flag.Duration("d", 10, "Timeout between ticks")
	flag.Parse()

	fmt.Println(*host, *walletAddress, *duration)

	for range time.Tick(*duration * time.Second) {
		go sendRequest(host, walletAddress)
	}

}

func sendRequest(host, walletAddress *string) {
	resp, err := http.Get(*host + "/" + *walletAddress)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
