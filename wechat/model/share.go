package model

//Share represents weixin share ticket
type Share struct {
	OpenGId   string    `json:"openGId"`
	Watermark Watermark `json:"watermark"`
}
