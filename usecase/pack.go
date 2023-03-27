package usecase

import (
	"playground/pipeline/entity"
)

type Pack struct{}

func NewPack() Pack {
	return Pack{}
}

func (p Pack) Packing(done <-chan struct{}, in <-chan entity.Phone) <-chan entity.Phone {
	out := make(chan entity.Phone)

	go func() {
		defer close(out)

		for data := range in {
			data.IsAlreadyPacked = true

			select {
			case out <- data:
			case <-done:
				return
			}
		}
	}()

	return out
}
