package usecase

import (
	"playground/pipeline/entity"
)

type Wrap struct{}

func NewWrap() Wrap {
	return Wrap{}
}

func (w Wrap) Wrapping(done <-chan struct{}, in <-chan entity.Phone) <-chan entity.Box {
	out := make(chan entity.Box)

	go func() {
		temp := []entity.Phone{}
		defer close(out)

		for data := range in {
			select {
			default:
				if len(temp) < entity.MaxUnitPerBox {
					temp = append(temp, data)
				} else {
					out <- temp

					// reset
					temp = append([]entity.Phone{}, data)
				}
			case <-done:
				return
			}
		}

		// flush remaining data
		out <- temp
	}()

	return out
}
