package usecase

import (
	"playground/pipeline/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPack(t *testing.T) {
	dep := NewPack()
	assert.NotNil(t, dep)
}

func TestPack_Packing_Success(t *testing.T) {
	units := 8
	p := NewPack()

	outChecking := make(chan entity.Phone)
	go func() {
		defer close(outChecking)
		for i := 0; i < units; i++ {
			outChecking <- entity.Phone{
				Numbering: i + 1,
			}
		}
	}()
	done := make(chan struct{})
	defer close(done)

	out := p.Packing(done, outChecking)

	l := 0
	for data := range out {
		if data.IsAlreadyPacked {
			l++
		}
	}

	assert.Equal(t, l, units)
}

func TestPack_Packing_SuccessDone(t *testing.T) {
	units := 1000000000
	p := NewPack()

	outChecking := make(chan entity.Phone)
	go func() {
		defer close(outChecking)
		for i := 0; i < units; i++ {
			outChecking <- entity.Phone{
				Numbering: i + 1,
			}
		}
	}()
	done := make(chan struct{})
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(done)
	}()

	out := p.Packing(done, outChecking)

	l := 0
	for data := range out {
		if data.IsAlreadyPacked {
			l++
		}
	}

	assert.NotEqual(t, l, units)
}
