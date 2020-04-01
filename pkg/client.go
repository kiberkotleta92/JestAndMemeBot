package pkg

import (
	"JestAndMemeBot/pkg/random"
	"JestAndMemeBot/pkg/translate"
	"JestAndMemeBot/pkg/vk"

	"github.com/valyala/fasthttp"
)

//Client is a core struct which concatinates all app logic
type Client struct {
	InternalClient        *fasthttp.Client
	VKToken               string
	RandomToken           string
	YandexToken           string
	TextContentURL        string
	PictureContentGroupID string
}

//GetPicture returns url of random picture from wall of client's vk group id
func (c *Client) GetPicture() string {
	n := random.GetRandomNumber(c.InternalClient, c.RandomToken)
	return vk.GetPictureURL(c.InternalClient, c.VKToken, c.PictureContentGroupID, n)
}

//GetPrediction returns translated text from given page and line in client's text content url
func (c *Client) GetPrediction(page, line int) string {
	return translate.GetPrediction(c.InternalClient, c.YandexToken, c.TextContentURL, page, line)
}
