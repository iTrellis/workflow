// MIT License
// Copyright (c) 2015 go-trellis

package workflow

// StepFunc step function 执行函数
type StepFunc func(context Context) error

// Step step to do something
type Step struct {
	Label     string
	Run       StepFunc
	IsLast    bool
	DependsOn []*Step
}

// FailureFunc 错误后执行函数
type FailureFunc func(err error, step *Step, context Context) error
