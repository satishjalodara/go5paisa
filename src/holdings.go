package go5paisa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Holding represents a single holding
type Holding struct {
	BseCode         int `json:"BseCode"`
	CurrentPrice    int `json:"CurrentPrice"`
	DPQty           int `json:"DPQty"`
	Exchange        int `json:"Exch"`
	ExchangeType    int `json:"ExchType"`
	Name            int `json:"FullName"`
	NseCode         int `json:"NseCode"`
	POASigned       int `json:"POASigned"`
	PoolQty         int `json:"PoolQty"`
	Quantity        int `json:"Quantity"`
	ScripMultiplier int `json:"ScripMultiplier"`
	Symbol          int `json:"Symbol"`
}

// Data has all holdings for a user
type responseData struct {
	Head interface{} `json:"head"`
	Body Holdings    `json:"body"`
}

type Holdings struct {
	Data []Holding `json:"Data"`
}

func parsHoldingsResponse(resBody []byte, obj Holdings) {
	var body responseData
	body.Body = obj
	if err := json.Unmarshal(resBody, &body); err != nil {
		log.Fatal("Error parsing JSON response:", err)
	}
}

// GetHoldings fetches holdings of the user
func (c *Client) GetHoldings() {
	c.appConfig.head.RequestCode = holdingsRequestCode
	payload := genericPayloadBody{
		ClientCode: c.clientCode,
	}
	jsonValue, _ := json.Marshal(payload)
	res, err := c.connection.Post(baseURL+holdingsRoute, contentType, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))
	var holdings Holdings
	parseResBody(resBody, &holdings)
	fmt.Println(holdings)
}
