package model

//Profile represents weixin user profile
type Profile struct {
	OpenID    string    `json:"openId"`
	UnionID   string    `json:"unionId"`
	NickName  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarURL string    `json:"avatarUrl"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark"`
}
