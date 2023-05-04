package models

type GiveMe struct {
	Day string `json:"day"`
}

type Answer struct {
	StaffName string  `json:"employe"`
	Category  string  `json:"category"`
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Summ      float32 `json:"summ"`
}
