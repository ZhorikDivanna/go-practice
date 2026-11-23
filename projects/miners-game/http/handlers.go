package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"project-miners/game"
	"time"
)

type GameHandlers struct {
	enterprise *game.Enterprice
}

func NewGameHandlers(e *game.Enterprice) *GameHandlers {
	return &GameHandlers{
		enterprise: e,
	}
}

func (h *GameHandlers) HandleHire(w http.ResponseWriter, r *http.Request) {
	var request HireMinerRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := request.ValidateForRequest(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.enterprise.HireMiner(request.MinerType); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *GameHandlers) HandleStatus(w http.ResponseWriter, r *http.Request) {
	balance, activeCount, equipment, totalHired := h.enterprise.GetStatus()

	status := EnterpriseStatusDTO{
		Balance:           int(balance),
		ActiveMinersCount: activeCount,
		Equipment:         equipment,
		TotalHired:        totalHired,
	}

	b, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		errDTO := ErrorDTO{
			Message: "Ошибка при формировании отчета",
			Time: time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *GameHandlers) HandleBuy(w http.ResponseWriter, r *http.Request) {
	var buy BuyEquipmentRequest

	if err := json.NewDecoder(r.Body).Decode(&buy); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := buy.ValidateForBuy(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.enterprise.BuyEquipment(buy.Name); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *GameHandlers) HandleFinish(w http.ResponseWriter, r *http.Request) {
	stats, err := h.enterprise.FinishGame()
	if err != nil {
		w.Write([]byte("Вы еще не готовы завершить игру"))
		return
	}

	b, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

	go func() {
		time.Sleep(1 *time.Second)
		fmt.Println("Игра Завершена. Спасибо за участие!")
		os.Exit(0)
	}()
}

func (h *GameHandlers) HandleMinerPrices(w http.ResponseWriter, r *http.Request) {
	prices := MinersPricesDTO{
		Small:  5,
		Normal: 50,
		Strong: 450,
	}

	b, _ := json.MarshalIndent(prices, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (h *GameHandlers) HandleEquipmentPrices(w http.ResponseWriter, r *http.Request) {
	_, _, boughtMap, _ := h.enterprise.GetStatus()

	prices := EquipmentPricesDTO{
		Items: []EquipmentItem{
			{Name: "pickaxe", Cost: 3000, IsBought: boughtMap["pickaxe"]},
			{Name: "ventilation", Cost: 15000, IsBought: boughtMap["ventilation"]},
			{Name: "trolley", Cost: 50000, IsBought: boughtMap["trolley"]},
		},
	}

	b, _ := json.MarshalIndent(prices, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (h *GameHandlers) HandleActiveMiners(w http.ResponseWriter, r *http.Request) {

	filterClass := r.URL.Query().Get("class")

	miners := h.enterprise.GetActiveMiners(filterClass)

	b, _ := json.MarshalIndent(miners, "", "    ")
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}