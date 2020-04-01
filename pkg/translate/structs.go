package translate

//YandexResponse is self explanatory, need it for fast (un)marshalling with easyjson
type YandexResponse struct {
	Code int64    `json:"code"`
	Lang string   `json:"lang"`
	Text []string `json:"text"`
}
