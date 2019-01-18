// MIT License

// Copyright (c) 2015 go-trellis

package workflow_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-trellis/workflow"
)

func TestWorkflow(t *testing.T) {
	var err error

	// right workflow
	go func() {
		w := workflow.New()
		step1 := &workflow.Step{
			Label: "one",
			Run:   stepOne,
		}

		step2 := &workflow.Step{
			Label:     "two",
			Run:       stepTwo,
			DependsOn: []*workflow.Step{step1},
		}

		step3 := &workflow.Step{
			Label:     "three",
			Run:       stepThree,
			IsLast:    true,
			DependsOn: []*workflow.Step{step2, step1},
		}
		err = w.SetStartStep(step3).
			SetFailureFunc(stepFailedFunc).
			SetLastStepConcurrency(true).
			Run(workflow.Context{})
	}()

	time.Sleep(time.Second * 3)
	if err != nil {
		t.Error(err)
	}

	// error workflow
	go func() {
		w := workflow.New()
		step1 := &workflow.Step{
			Label: "one",
			Run:   stepOne,
		}

		step2 := &workflow.Step{
			Label:     "two",
			Run:       stepTwoFailed,
			DependsOn: []*workflow.Step{step1},
		}

		step3 := &workflow.Step{
			Label:     "three",
			Run:       stepThree,
			IsLast:    true,
			DependsOn: []*workflow.Step{step2, step1},
		}
		err = w.SetStartStep(step3).
			SetFailureFunc(stepFailedFunc).
			SetLastStepConcurrency(true).
			Run(workflow.Context{})
	}()

	time.Sleep(time.Second * 3)
	if err == nil {
		t.Error("failed get workflow error")
	}
}

func stepThree(context workflow.Context) error {
	fmt.Println("step 3")
	context["text3"] = "zzz"
	time.Sleep(time.Second * 2)
	fmt.Println(context)
	fmt.Println("step 3: self finish")
	return nil
}

func stepTwo(context workflow.Context) error {
	fmt.Println("step 2")
	context["text2"] = "yyy"
	fmt.Println(context)
	return nil
}

func stepTwoFailed(context workflow.Context) error {
	fmt.Println("step 2")
	context["text2"] = "yyy"
	fmt.Println(context)
	return fmt.Errorf("failed 2")
}

func stepOne(context workflow.Context) error {
	fmt.Println("step 1")
	context["text1"] = "xxx"
	fmt.Println(context)
	return nil
}

func stepFailedFunc(err error, step *workflow.Step, context workflow.Context) error {
	fmt.Println(err)
	fmt.Println(*step)
	fmt.Println(context)
	return nil
}
