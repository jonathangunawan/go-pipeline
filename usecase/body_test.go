package usecase

import (
	"playground/pipeline/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBody(t *testing.T) {
	dep := NewBody()
	assert.NotNil(t, dep)
}

func TestBody_BodyAndSerialize_Success(t *testing.T) {
	units := 8
	b := NewBody()

	outPlan := make(chan entity.Phone)
	go func() {
		defer close(outPlan)
		for i := 0; i < units; i++ {
			outPlan <- entity.Phone{}
		}
	}()
	done := make(chan struct{})
	defer close(done)

	out := b.BodyAndSerialize(done, outPlan)

	l := 0
	for data := range out {
		if data.SerialNumber != "" {
			l++
		}
	}

	assert.Equal(t, l, units)
}

func TestBody_BodyAndSerialize_SuccessDone(t *testing.T) {
	units := 1000000000
	b := NewBody()

	outPlan := make(chan entity.Phone)
	go func() {
		defer close(outPlan)
		for i := 0; i < units; i++ {
			outPlan <- entity.Phone{
				Numbering: i + 1,
			}
		}
	}()
	done := make(chan struct{})
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(done)
	}()

	out := b.BodyAndSerialize(done, outPlan)

	l := 0
	for data := range out {
		if data.SerialNumber != "" {
			l++
		}
	}

	assert.NotEqual(t, l, units)
}
