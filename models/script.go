package models

type Script struct {
	CurrentPrice float64 `json:"c"`
	High         float64 `json:"h"`
	Low          float64 `json:"l"`
	Open         float64 `json:"o"`
	PrevClose    float64 `json:"pc"`
	Timestamp    int64   `json:"t"`
}
