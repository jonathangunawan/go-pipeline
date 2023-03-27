package usecase

import (
	"playground/pipeline/entity"

	"github.com/google/uuid"
)

type Body struct {
}

func NewBody() Body {
	return Body{}
}

func (b Body) BodyAndSerialize(done <-chan struct{}, in <-chan entity.Phone) <-chan entity.Phone {
	out := make(chan entity.Phone)

	go func() {
		defer close(out)

		for data := range in {
			data.SerialNumber = uuid.NewString()

			select {
			case out <- data:
			case <-done:
				return
			}
		}
	}()

	return out
}
