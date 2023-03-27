package usecase

import (
	"playground/pipeline/entity"
)

type Plan struct{}

func NewPlan() Plan {
	return Plan{}
}

func (p Plan) PlanProduct(done <-chan struct{}, units int) <-chan entity.Phone {
	out := make(chan entity.Phone)
	go func() {
		defer close(out)
		for i := 0; i < units; i++ {
			data := entity.Phone{
				Numbering: i + 1,
			}

			select {
			case out <- data:
			case <-done:
				return
			}

		}
	}()

	return out
}
