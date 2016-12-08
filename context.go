// MIT License

// Copyright (c) 2015 rutcode-go

package workflow

import (
	"fmt"
)

type Context map[string]interface{}

func (p *Workflow) GetContextString(key string) (val string, err error) {
	return p.context.GetContextString(key)
}

func (p *Workflow) GetContextInt(key string) (val int, err error) {
	return p.context.GetContextInt(key)
}

func (p *Workflow) GetContextBool(key string) (val bool, err error) {
	return p.context.GetContextBool(key)
}

func (p *Workflow) GetContextInterface(key string) (val interface{}, err error) {
	return p.context.GetContextInterface(key)
}

func (p Context) GetContextString(key string) (val string, err error) {
	if p == nil {
		err = fmt.Errorf("context is empty")
		return
	}
	if v, e := p[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	} else if sv, ok := v.(string); ok {
		val = sv
		return
	}

	err = fmt.Errorf("value of %s's type is not string", key)
	return
}

func (p Context) GetContextInt(key string) (val int, err error) {
	if p == nil {
		err = fmt.Errorf("context is empty")
		return
	}
	if v, e := p[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	} else if iv, ok := v.(int); ok {
		val = iv
		return
	}

	err = fmt.Errorf("value of %s's type is not string", key)
	return
}

func (p Context) GetContextBool(key string) (val bool, err error) {
	if p == nil {
		err = fmt.Errorf("context is empty")
		return
	}
	if v, e := p[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	} else if iv, ok := v.(bool); ok {
		val = iv
		return
	}

	err = fmt.Errorf("value of %s's type is not bool", key)
	return
}

func (p Context) GetContextInterface(key string) (val interface{}, err error) {
	if p == nil {
		err = fmt.Errorf("context is empty")
		return
	}
	var e bool
	if val, e = p[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	}

	return
}
