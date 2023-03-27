package main

import (
	"fmt"
	"playground/pipeline/entity"
	"playground/pipeline/usecase"
	"time"
)

func main() {
	now := time.Now()
	input := 8
	laneNumbers := 5
	isMultiLane := true

	done := make(chan struct{})
	defer close(done)

	pln := usecase.NewPlan()
	bdy := usecase.NewBody()
	ins := usecase.NewInstall()
	chk := usecase.NewCheck()
	mltchk := usecase.NewMulti()
	pck := usecase.NewPack()
	wrp := usecase.NewWrap()

	outPlan := pln.PlanProduct(done, input)
	outBody := bdy.BodyAndSerialize(done, outPlan)
	outInstall := ins.InstallDisplayEtc(done, outBody)

	outChecking := make(<-chan entity.Phone)
	if !isMultiLane {
		outChecking = chk.CheckFunctionality(done, outInstall)
	} else {
		outChecking = mltchk.MultiLaneCheck(laneNumbers, done, outInstall)
	}

	outPack := pck.Packing(done, outChecking)
	outWrap := wrp.Wrapping(done, outPack)

	boxes := []entity.Box{}
	for n := range outWrap {
		boxes = append(boxes, n)
		// fmt.Println(n)
	}

	fmt.Println("Total Box: ", len(boxes))
	fmt.Println("Duration: ", time.Since(now))
}

// normal pipeline | 8 inputs, max 5 units per box: ~16s
// fanout fanin pipeline | 8 goroutine | 8 inputs, max 5 units per box: ~2s
