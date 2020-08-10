package go5paisa

import "encoding/json"
import "log"

type body struct {
	ClientCode string `json:"ClientCode"`
	Message    string `json:"Message"`
}

type responseBody struct {
	Head interface{} `json:"head"`
	Body interface{} `json:"body"`
}

func parseResBody(resBody []byte, obj interface{}) {
	var body responseBody
	body.Body = obj
	if err := json.Unmarshal(resBody, &body); err != nil {
		log.Fatal("Error parsing JSON response:", err)
	}
}
