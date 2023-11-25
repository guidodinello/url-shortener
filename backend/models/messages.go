package models

type RequestData struct {
	Url string `json:"url"`
}

type ResponseData struct {
	Key string `json:"shortUrl"`
	Url string `json:"url"`
}
