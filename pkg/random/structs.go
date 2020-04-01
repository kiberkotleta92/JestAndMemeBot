package random

//RandomRequest is self explanatory, need it for fast (un)marshalling with easyjson
type RandomRequest struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		APIKey      string `json:"apiKey"`
		Max         int64  `json:"max"`
		Min         int64  `json:"min"`
		N           int64  `json:"n"`
		Replacement bool   `json:"replacement"`
	} `json:"params"`
}

//RandomResponse is self explanatory, need it for fast (un)marshalling with easyjson
type RandomResponse struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AdvisoryDelay int64 `json:"advisoryDelay"`
		BitsLeft      int64 `json:"bitsLeft"`
		BitsUsed      int64 `json:"bitsUsed"`
		Random        struct {
			CompletionTime string  `json:"completionTime"`
			Data           []int64 `json:"data"`
		} `json:"random"`
		RequestsLeft int64 `json:"requestsLeft"`
	} `json:"result"`
}
