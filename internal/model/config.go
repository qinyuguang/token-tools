package model

type TokenConfig struct {
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Network string `json:"network"`
}

type AccountConfig struct {
	Name    string `json:"name"`
	Alias   string `json:"alias"`
	Address string `json:"address"`
}
