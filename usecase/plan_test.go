package usecase

import (
	"playground/pipeline/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPlan(t *testing.T) {
	dep := NewPlan()
	assert.NotNil(t, dep)
}

func TestPlan_PlanProduct_Success(t *testing.T) {
	units := 8
	p := NewPlan()
	done := make(chan struct{})
	defer close(done)

	out := p.PlanProduct(done, units)
	l := 0
	def := entity.Phone{}

	for data := range out {

		if data != def {
			l++
		}
	}

	assert.Equal(t, l, units)
}

func TestPlan_PlanProduct_SuccessDone(t *testing.T) {
	units := 1000000000
	p := NewPlan()
	done := make(chan struct{})
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(done)
	}()

	out := p.PlanProduct(done, 1000000000)
	l := 0
	def := entity.Phone{}

	for data := range out {
		if data != def {
			l++
		}
	}

	assert.NotEqual(t, l, units)
}
