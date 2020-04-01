package random

import (
	"github.com/valyala/fasthttp"
)

//GetRandomNumber asks random.org for an integer from 0 to 145
func GetRandomNumber(client *fasthttp.Client, token string) int64 {
	var responseStruct RandomResponse

	requestStruct := RandomRequest{
		Jsonrpc: "2.0",
		Method:  "generateIntegers",
		Params: struct {
			APIKey      string "json:\"apiKey\""
			Max         int64  "json:\"max\""
			Min         int64  "json:\"min\""
			N           int64  "json:\"n\""
			Replacement bool   "json:\"replacement\""
		}{
			APIKey:      token,
			Max:         145,
			Min:         0,
			N:           1,
			Replacement: true,
		},
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	requestBody, _ := requestStruct.MarshalJSON()

	req.SetRequestURI("https://api.random.org/json-rpc/1/invoke")
	req.SetBody(requestBody)

	client.Do(req, resp)

	bodyBytes := resp.Body()
	responseStruct.UnmarshalJSON(bodyBytes)

	return responseStruct.Result.Random.Data[0]
}
