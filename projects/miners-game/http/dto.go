package http

import (
	"encoding/json"
	"errors"
	"time"
)

type HireMinerRequestDTO struct {
	MinerType string `json:"miner_type"`
}

type BuyEquipmentRequest struct {
	Name string `json:"name"`
}

type EnterpriseStatusDTO struct {
	Balance           int             `json:"balance"`
	ActiveMinersCount int             `json:"active_miners"`
	Equipment         map[string]bool `json:"equipment"`
	TotalHired        map[string]int  `json:"total_hired"`
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

type MinersPricesDTO struct {
	Small  int `json:"small_miner_cost"`
	Normal int `json:"normal_miner_cost"`
	Strong int `json:"strong_miner_cost"`
}

type EquipmentItem struct {
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	IsBought bool   `json:"is_bought"`
}

type EquipmentPricesDTO struct {
	Items []EquipmentItem `json:"items"`
}

func (h HireMinerRequestDTO) ValidateForRequest() error {
	if h.MinerType == "" {
		return errors.New("Пустой ввод")
	}
	return nil
}

func (b BuyEquipmentRequest) ValidateForBuy() error {
	if b.Name == "" {
		return errors.New("Пустой ввод")
	}
	return nil
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
