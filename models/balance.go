package models

type Invoke struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int32  `json:"id"`
	Result  Result `json:"result"`
}

type Result struct {
	Script       string  `json:"script"`
	State        string  `json:"state"`
	Gas_consumed string  `json:"gas_consumed"`
	Stack        []Stack `json:"stack"`
}

type Stack struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

/*
{
    "jsonrpc": "2.0",
    "id": 3,
    "result": {
        "script": "142b41aea9d405fef2e809e3c8085221ce944527a751c10962616c616e63654f6667f91d6b7085db7c5aaf09f19eeec1ca3c0db2c6ec",
        "state": "HALT, BREAK",
        "gas_consumed": "0.338",
        "stack": [
            {
                "type": "ByteArray",
                "value": "0084d717"
            }
        ]
    }
}
*/
