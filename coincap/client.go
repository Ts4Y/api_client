package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {

	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero")
	}

	return &Client{
		&http.Client{
			Timeout: timeout,
		},
	}, nil

}

func (c Client) GetCoinInfo(name string) (Coin, error) {

	url := fmt.Sprintf("https://api.coincap.io/v2/assets/ %s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return Coin{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var a CoinResponse
	if err = json.Unmarshal(body, &a); err != nil {
		return Coin{}, err
	}

	return a.Coin, nil
}

func (c Client) GetCoinMarketInfo (baseID string)(MarketsList, error){
	marketsUrl := fmt.Sprintf("https://api.coincap.io/v2/assets/%s/markets",baseID)
	resp,err := c.client.Get(marketsUrl)
	if err != nil{
		return MarketsList{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln("не удалось прочитать тело запроса, ошибка:",err)
	}

	var ml MarketsList
	if err = json.Unmarshal(body, &ml); err != nil{
		return MarketsList{},nil
	}

	return ml, nil
}
