package usecase

import (
	"playground/pipeline/entity"
	"time"
)

type Check struct{}

func NewCheck() Check {
	return Check{}
}

func (c Check) CheckFunctionality(done <-chan struct{}, in <-chan entity.Phone) <-chan entity.Phone {
	out := make(chan entity.Phone)

	go func() {
		defer close(out)

		for data := range in {
			time.Sleep(2 * time.Second)
			data.IsFunctional = true

			select {
			case out <- data:
			case <-done:
				return
			}
		}
	}()

	return out
}
