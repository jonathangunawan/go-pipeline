package usecase

import "playground/pipeline/entity"

type Install struct{}

func NewInstall() Install {
	return Install{}
}

func (i Install) InstallDisplayEtc(done <-chan struct{}, in <-chan entity.Phone) <-chan entity.Phone {
	out := make(chan entity.Phone)

	go func() {
		defer close(out)
		for data := range in {
			data.IsBatteryInstalled = true
			data.IsDisplayInstalled = true

			select {
			case out <- data:
			case <-done:
				return
			}
		}
	}()

	return out
}
