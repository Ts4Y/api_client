package main

import (
	"api_client/coincap"
	"fmt"
	"log"
	"time"
)

func main() {

	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	coin, err := coincapClient.GetCoinInfo("bitcoin")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(coin.Info())

	// market,err := coincapClient.GetCoinMarketInfo("dogecoin")
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// for _,m:=range market.Market{
	// 	fmt.Println(m.MarketsInfo())
	// }
}
