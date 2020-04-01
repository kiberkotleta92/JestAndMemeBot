package vk

//VKResponse is self explanatory, need it for fast (un)marshalling with easyjson
type VKResponse struct {
	Response struct {
		Count int64 `json:"count"`
		Items []struct {
			AlbumID int64 `json:"album_id"`
			Date    int64 `json:"date"`
			ID      int64 `json:"id"`
			OwnerID int64 `json:"owner_id"`
			PostID  int64 `json:"post_id"`
			Sizes   []struct {
				Height int64  `json:"height"`
				Type   string `json:"type"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"sizes"`
			Text   string `json:"text"`
			UserID int64  `json:"user_id"`
		} `json:"items"`
	} `json:"response"`
}
