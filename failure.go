// MIT License

// Copyright (c) 2015 rutcode-go

package workflow

type FailureFunc func(err error, step *Step, context Context) error
