package game

import (
	"context"
	"time"
)

type Coal int

type MinerInfo struct {
	Class  string `json:"class"`
	Energy int    `json:"energy"`
}

type Miner interface {
	Run(ctx context.Context) <-chan Coal
	Info() MinerInfo
}

type SmallMiner struct {
	Energy int
}

func NewSmallMiner() *SmallMiner {
	return &SmallMiner{
		Energy: 30,
	}
}

func (m *SmallMiner) Info() MinerInfo {
	return MinerInfo{
		Class:  "small",
		Energy: m.Energy,
	}
}

func (m *SmallMiner) Run(ctx context.Context) <-chan Coal {
	ch := make((chan Coal))

	go func() {
		defer close(ch)

		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(3 * time.Second):
				ch <- 1
				m.Energy--
			}
		}
	}()
	return ch
}

type NormalMiner struct {
	Energy int
}

func NewNormalMiner() *NormalMiner {
	return &NormalMiner{
		Energy: 45,
	}
}

func (m *NormalMiner) Info() MinerInfo {
	return MinerInfo{
		Class:  "normal",
		Energy: m.Energy,
	}
}

func (m *NormalMiner) Run(ctx context.Context) <-chan Coal {
	ch := make((chan Coal))

	go func() {
		defer close(ch)

		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(2 * time.Second):
				ch <- 3
				m.Energy--
			}
		}
	}()
	return ch
}

type StrongMiner struct {
	Energy     int
	MiningRate int
}

func NewStrongMiner() *StrongMiner {
	return &StrongMiner{
		Energy:     60,
		MiningRate: 10,
	}
}

func (m *StrongMiner) Info() MinerInfo {
	return MinerInfo{
		Class:  "strong",
		Energy: m.Energy,
	}
}

func (m *StrongMiner) Run(ctx context.Context) <-chan Coal {
	ch := make((chan Coal))

	go func() {
		defer close(ch)

		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):
				ch <- Coal(m.MiningRate)
				m.MiningRate += 3

				m.Energy--
			}
		}
	}()
	return ch
}
