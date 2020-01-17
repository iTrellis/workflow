// MIT License
// Copyright (c) 2015 go-trellis

package workflow

// Workflow work flow
type Workflow struct {
	start *Step
	// context     Context
	failureFunc FailureFunc

	queue   []*Step
	inQueue map[*Step]bool

	lastStepConcurrency bool
}

// New 生成workflow对象
func New() *Workflow {
	p := &Workflow{}
	p.inQueue = make(map[*Step]bool)
	// p.context = make(map[string]interface{})
	return p
}

// NewWithFailedFunc 生成包含错误执行方法的工作流对象
func NewWithFailedFunc(fun FailureFunc) *Workflow {
	p := New()
	p.failureFunc = fun
	return p
}

// Run 启动工作流
func (p *Workflow) Run(context Context) error {
	p.loadQueue(p.start)
	for _, step := range p.queue {
		if p.lastStepConcurrency && step.IsLast {
			go func() { _ = p.doRunStep(step, context) }()
			return nil
		}
		if err := p.doRunStep(step, context); err != nil {
			return err
		}
	}
	return nil
}

func (p *Workflow) doRunStep(step *Step, context Context) error {
	if err := step.Run(context); err != nil {
		if e := p.doFailureFunc(err, step, context); e != nil {
			err = e
		}
		return err
	}
	return nil
}

func (p *Workflow) doFailureFunc(err error, step *Step, context Context) (e error) {
	if p.failureFunc != nil {
		e = p.failureFunc(err, step, context)
	}
	return
}

// SetFailureFunc set failure function
// 设置错误后的执行方法
func (p *Workflow) SetFailureFunc(fun FailureFunc) *Workflow {
	p.failureFunc = fun
	return p
}

// SetStartStep set start step
// 设置第一个启动函数
func (p *Workflow) SetStartStep(s *Step) *Workflow {
	p.start = s
	return p
}

// SetLastStepConcurrency set the return flag if running before last step
// 在最后一步之前是否直接返回
func (p *Workflow) SetLastStepConcurrency(c bool) *Workflow {
	p.lastStepConcurrency = c
	return p
}

func (p *Workflow) loadQueue(s *Step) {
	if s == nil {
		return
	}

	for _, step := range s.DependsOn {
		p.loadQueue(step)
	}

	if !p.inQueue[s] {
		p.inQueue[s] = true
		p.queue = append(p.queue, s)
	}
}
