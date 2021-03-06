package block

import (
	"Neo-Rank/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Block struct {
}

func (b *Block) GetNep5Balance(host string, assetId string, address string) (models.Invoke, error) {

	/*
		payload := make(map[string]interface{})

		payload["jsonrpc"] = "2.0"
		payload["method"] = "invokefunction"

		var address = []map[string]string{
			{
				"type":  "Hash160",
				"value": "0xa7274594ce215208c8e309e8f2fe05d4a9ae412b",
			},
		}

		// payload["params"] =
		var params = []interface{}{
			"0xecc6b20d3ccac1ee9ef109af5a7cdb85706b1df9",
			"balanceOf",
			address,
		}

		payload["params"] = params

	*/

	data := fmt.Sprintf("{\"jsonrpc\": \"2.0\",\"method\": \"invokefunction\",  \"params\": [    \"%s\",    \"balanceOf\",    [      {\"type\": \"Hash160\",        \"value\": \"%s\"      }    ]  ],  \"id\": 3}", assetId, address)

	// fmt.Println("host", host)
	payload := strings.NewReader(data)

	req, _ := http.NewRequest("POST", host, payload)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	// fmt.Println("body", string(body))

	var invoke models.Invoke
	json.Unmarshal(body, &invoke)
	return invoke, err
}

func (b *Block) GetNep5Decimals(host string, assetId string) (models.Invoke, error) {

	data := fmt.Sprintf("{\"jsonrpc\": \"2.0\",\"method\": \"invokefunction\",  \"params\": [    \"%s\",    \"decimals\",    [    ]  ],  \"id\": 2}", assetId)

	// fmt.Println("data", data)
	payload := strings.NewReader(data)

	req, _ := http.NewRequest("POST", host, payload)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	// fmt.Println("body", string(body))

	var invoke models.Invoke
	json.Unmarshal(body, &invoke)
	return invoke, err
}
