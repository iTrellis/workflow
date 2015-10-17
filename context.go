package workflow

import (
	"fmt"
)

type Context map[string]interface{}

func (p *Workflow) GetContextString(key string) (val string, err error) {
	if v, e := p.context[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	} else if sv, ok := v.(string); ok {
		val = sv
		return
	}

	err = fmt.Errorf("value of %s's type is not string", key)
	return
}

func (p *Workflow) GetContextInt(key string) (val int, err error) {
	if v, e := p.context[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	} else if iv, ok := v.(int); ok {
		val = iv
		return
	}

	err = fmt.Errorf("value of %s's type is not string", key)
	return
}

func (p *Workflow) GetContextInterface(key string) (val interface{}, err error) {
	var e bool
	if val, e = p.context[key]; !e {
		err = fmt.Errorf("value of %s not in Context", key)
		return
	}

	return
}
