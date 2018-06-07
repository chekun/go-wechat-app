package model

//Run represents weixin run data
type Run struct {
	Data []RunRecord `json:"stepInfoList"`
}

//RunRecord represents weixin run data
type RunRecord struct {
	Timestamp uint `json:"timestamp"`
	Steps     uint `json:"step"`
}
