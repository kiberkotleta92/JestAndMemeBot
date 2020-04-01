package translate

import (
	"bytes"
	"net/url"

	"github.com/valyala/fasthttp"
)

//getText downloads text content of contentURL and returns line on a given page
func getText(client *fasthttp.Client, contentURL string, page, line int) string {
	storage := make([]byte, 4e+6)
	lineBytes := make([]byte, 100)
	lineCoord := page*26 + line
	_, bodyBytes, _ := client.Get(storage, contentURL)
	text := bytes.NewBuffer(bodyBytes)
	for i := 0; i < 316+lineCoord; i++ {
		lineBytes, _ = text.ReadBytes(byte('.'))
	}
	return string(lineBytes)
}

//getTranslation asks translate.yandex.ru to translate given text
func getTranslation(client *fasthttp.Client, token string, text string) string {
	var content YandexResponse
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	u, _ := url.Parse("https://translate.yandex.net/api/v1.5/tr.json/translate")
	q := url.Values{}
	q.Add("key", token)
	q.Add("text", text)
	q.Add("lang", "en-ru")
	q.Add("format", "plain")
	u.RawQuery = q.Encode()

	req.SetRequestURI(u.String())

	client.Do(req, resp)

	bodyBytes := resp.Body()
	content.UnmarshalJSON(bodyBytes)

	return content.Text[0]
}

//GetPrediction downloads text content of contentURL, takes line on a given page,
//asks translate.yandex.ru to translate line, returns translated line
func GetPrediction(client *fasthttp.Client, token string, contentURL string, page, line int) string {
	text := getText(client, contentURL, page, line)
	return getTranslation(client, token, text)
}
