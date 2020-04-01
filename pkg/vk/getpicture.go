package vk

import (
	"net/url"

	"github.com/valyala/fasthttp"
)

//GetPictureURL asks vk.com for 146 pictures (becase i remember memes from 2011 http://lurkmore.to/Чуров)
//from wall of given group id and takes one with given index.
func GetPictureURL(client *fasthttp.Client, token string, groupID string, randomIndex int64) string {
	var content VKResponse

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	u, _ := url.Parse("https://api.vk.com/method/photos.get")
	q := url.Values{}
	q.Add("access_token", token)
	q.Add("owner_id", groupID)
	q.Add("album_id", "wall")
	q.Add("count", "146")
	q.Add("v", "5.103")
	u.RawQuery = q.Encode()

	req.SetRequestURI(u.String())

	client.Do(req, resp)

	bodyBytes := resp.Body()
	content.UnmarshalJSON(bodyBytes)

	return content.Response.Items[randomIndex].Sizes[6].URL
}
