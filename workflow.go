package workflow

type Workflow struct {
	start       *Step
	context     Context
	failureFunc FailureFunc

	queue   []*Step
	inQueue map[*Step]bool

	lastStepConcurrency bool
}

func New() *Workflow {
	p := &Workflow{}
	p.inQueue = make(map[*Step]bool)
	p.context = make(map[string]interface{})
	return p
}

func NewWithFailedFunc(fun FailureFunc) *Workflow {
	p := New()
	p.failureFunc = fun
	return p
}

func (p *Workflow) Run() error {
	p.loadQueue(p.start)
	for _, step := range p.queue {
		if p.lastStepConcurrency && step.IsLast {
			go p.doRunStep(step)
			return nil
		}
		if err := p.doRunStep(step); err != nil {
			return err
		}
	}
	return nil
}

func (p *Workflow) doRunStep(step *Step) error {
	if err := step.Run(p.context); err != nil {
		if e := p.doFailureFunc(err, step); e != nil {
			err = e
		}
		return err
	}
	return nil
}

func (p *Workflow) SetContext(key string, val interface{}) *Workflow {
	if p.context == nil {
		p.context = make(map[string]interface{})
	}

	p.context[key] = val
	return p
}

func (p *Workflow) doFailureFunc(err error, step *Step) (e error) {
	if p.failureFunc != nil {
		e = p.failureFunc(err, step, p.context)
	}
	return
}

func (p *Workflow) SetFailureFunc(fun FailureFunc) *Workflow {
	p.failureFunc = fun
	return p
}

func (p *Workflow) SetStartStep(s *Step) *Workflow {
	p.start = s
	return p
}

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
	return
}
