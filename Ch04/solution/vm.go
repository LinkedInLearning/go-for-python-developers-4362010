package main

import (
	"fmt"
	"time"
)

const PerHour = 5 // ¢

type VM struct {
	StartTime time.Time
	EndTime   time.Time
}

func (v VM) Cost() float64 {
	end := v.EndTime
	if end.Equal(time.Time{}) {
		end = time.Now()
	}

	d := float64(end.Sub(v.StartTime)) / float64(time.Hour)
	return d * PerHour
}

type Spot struct {
	VM
}

func (s Spot) Cost() float64 {
	p := s.VM.Cost()
	return p * 0.8
}

type Coster interface {
	Cost() float64
}

func TotalCost(vms []Coster) float64 {
	total := 0.0
	for _, vm := range vms {
		total += vm.Cost()
	}
	return total
}

func main() {
	vms := []Coster{
		&VM{
			StartTime: time.Date(2022, time.April, 12, 17, 30, 0, 0, time.UTC),
			EndTime:   time.Date(2022, time.April, 12, 19, 54, 0, 0, time.UTC),
		},
		Spot{
			VM: VM{
				StartTime: time.Date(2022, time.April, 13, 9, 46, 0, 0, time.UTC),
				EndTime:   time.Date(2022, time.April, 15, 12, 7, 0, 0, time.UTC),
			},
		},
	}
	fmt.Println(TotalCost(vms))
}
