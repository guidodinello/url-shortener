package models

type RequestData struct {
	Url string `json:"url"`
}

type ResponseData struct {
	ShortUrl string `json:"shortUrl"`
	Url      string `json:"url"`
}
