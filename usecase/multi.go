package usecase

import (
	"playground/pipeline/entity"
	"sync"
)

type Multi struct{}

func NewMulti() Multi {
	return Multi{}
}

func (m Multi) MultiLaneCheck(laneNumbers int, done <-chan struct{}, in <-chan entity.Phone) <-chan entity.Phone {
	var wg sync.WaitGroup
	out := make(chan entity.Phone)

	lane := laneNumbers
	temp := []<-chan entity.Phone{}

	for i := 0; i < lane; i++ {
		chk := NewCheck()
		temp = append(temp, chk.CheckFunctionality(done, in))
	}

	output := func(c <-chan entity.Phone) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(lane)
	for _, c := range temp {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
