package go5paisa

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Holding represents a single holding
type Holding struct {
	BseCode         int     `json:"BseCode"`
	CurrentPrice    float32 `json:"CurrentPrice"`
	DPQty           int     `json:"DPQty"`
	Exchange        string  `json:"Exch"`
	ExchangeType    string  `json:"ExchType"`
	Name            string  `json:"FullName"`
	NseCode         int     `json:"NseCode"`
	POASigned       string  `json:"POASigned"`
	PoolQty         int     `json:"PoolQty"`
	Quantity        int     `json:"Quantity"`
	ScripMultiplier int     `json:"ScripMultiplier"`
	Symbol          string  `json:"Symbol"`
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
func (c *Client) GetHoldings() (Holdings, error) {
	var holdings Holdings
	c.appConfig.head.RequestCode = holdingsRequestCode
	payloadBody := genericPayloadBody{
		ClientCode: c.clientCode,
	}
	payload := genericPayload{
		Head: c.appConfig.head,
		Body: payloadBody,
	}
	jsonValue, _ := json.Marshal(payload)
	res, err := c.connection.Post(baseURL+holdingsRoute, contentType, bytes.NewBuffer(jsonValue))
	if err != nil {
		return holdings, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return holdings, err
	}

	parseResBody(resBody, &holdings)
	return holdings, nil
}
