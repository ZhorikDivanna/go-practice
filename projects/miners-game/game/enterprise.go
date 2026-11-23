package game

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"sync"
	"time"
)

type GameStats struct {
	TotalTime    float64        `json:"total_time_seconds"`
	TotalHired   map[string]int `json:"total_hired"`
	FinalBalance Coal           `json:"final_balance"`
}

type Enterprice struct {
	Balance         Coal
	ActiveMiners    []Miner
	TotalMiners     map[string]int
	mtx             sync.RWMutex
	ctx             context.Context
	cancel          context.CancelFunc
	BoughtEquipment map[string]bool
	StartTime       time.Time
}

func NewEnterPrise() *Enterprice {
	ctx, cancel := context.WithCancel(context.Background())

	e := &Enterprice{
		Balance:         0,
		ActiveMiners:    make([]Miner, 0),
		TotalMiners:     make(map[string]int),
		ctx:             ctx,
		cancel:          cancel,
		BoughtEquipment: make(map[string]bool),
		StartTime:       time.Now(),
	}

	go func() {
		for {
			select {
			case <-e.ctx.Done():
				return
			case <-time.After(1 * time.Second):

				e.mtx.Lock()
				e.Balance++
				e.mtx.Unlock()
			}
		}
	}()
	return e
}

func (e *Enterprice) HireMiner(minerType string) error {
	var cost Coal
	var m Miner

	switch minerType {
	case "small":
		cost = 5
		m = NewSmallMiner()
	case "normal":
		cost = 50
		m = NewNormalMiner()
	case "strong":
		cost = 450
		m = NewStrongMiner()
	default:
		return errors.New("ошибка типа шахтера")
	}

	e.mtx.Lock()
	if e.Balance < cost {
		e.mtx.Unlock()
		return fmt.Errorf("недостаточно денег для %s. Нужно: %d, есть: %d", minerType, cost, e.Balance)
	}

	e.Balance -= cost
	e.ActiveMiners = append(e.ActiveMiners, m)
	e.TotalMiners[minerType]++

	e.mtx.Unlock()

	coalChan := m.Run(e.ctx)

	go func() {
		for v := range coalChan {
			e.mtx.Lock()
			e.Balance += v
			e.mtx.Unlock()
		}
		e.RemoveMiner(m)

	}()

	return nil
}

func (e *Enterprice) BuyEquipment(name string) error {
	var cost Coal = 3_000

	switch name {
	case "pickaxe":
		cost = 3_000
	case "ventilation":
		cost = 15_000
	case "trolley":
		cost = 50_000
	default:
		return errors.New("Неизветное оборудование")
	}

	e.mtx.Lock()
	defer e.mtx.Unlock()

	if e.BoughtEquipment[name] {
		return errors.New("Этот предмет уже куплен")
	}

	if e.Balance < cost {
		return fmt.Errorf("недостаточно денег чтобы купить %s", name)
	}

	e.Balance -= cost
	e.BoughtEquipment[name] = true

	return nil
}

func (e *Enterprice) FinishGame() (GameStats, error) {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	if len(e.BoughtEquipment) < 3 {
		return GameStats{}, errors.New("Вы не купили все необходимые предметы для завершения игры")
	}

	e.cancel()

	stats := GameStats{
		TotalTime:    time.Since(e.StartTime).Seconds(),
		TotalHired:   maps.Clone(e.TotalMiners),
		FinalBalance: e.Balance,
	}

	return stats, nil

}

func (e *Enterprice) GetStatus() (Coal, int, map[string]bool, map[string]int) {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	return e.Balance, len(e.ActiveMiners), maps.Clone(e.BoughtEquipment), maps.Clone(e.TotalMiners)
}

func (e *Enterprice) RemoveMiner(m Miner) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	for i, activeMiner := range e.ActiveMiners {
		if activeMiner == m {
			e.ActiveMiners = append(e.ActiveMiners[:i], e.ActiveMiners[i+1:]...)
			break
		}
	}
}

func (e *Enterprice) GetActiveMiners(filterClass string) []MinerInfo {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	var result []MinerInfo

	for _, m := range e.ActiveMiners {
		info := m.Info() // Вызываем метод интерфейса

		if filterClass == "" || info.Class == filterClass {
			result = append(result, info)
		}
	}

	return result
}
