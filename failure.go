package workflow

type FailureFunc func(err error, step *Step, context Context) error
