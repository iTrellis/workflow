// MIT License

// Copyright (c) 2015 go-trellis

package workflow

import (
	"fmt"
)

// Context 上下文
type Context map[string]interface{}

// GetContextString get a string value by key from context
func (p Context) GetContextString(key string) (val string, err error) {

	v, e := p.GetContextInterface(key)
	if e != nil {
		return "", e
	}
	if iv, ok := v.(string); ok {
		return iv, nil
	}

	return "", fmt.Errorf("value of %s's type is not string", key)
}

// GetContextInt get a int value by key from context
func (p Context) GetContextInt(key string) (val int, err error) {

	v, e := p.GetContextInterface(key)
	if e != nil {
		return 0, e
	}
	if iv, ok := v.(int); ok {
		return iv, nil
	}

	return 0, fmt.Errorf("value of %s's type is not int", key)
}

// GetContextBool get a bool value by key from context
func (p Context) GetContextBool(key string) (bool, error) {

	v, e := p.GetContextInterface(key)
	if e != nil {
		return false, e
	}
	if iv, ok := v.(bool); ok {
		return iv, nil
	}

	return false, fmt.Errorf("value of %s's type is not bool", key)
}

// GetContextInterface get context value
func (p Context) GetContextInterface(key string) (interface{}, error) {
	if p == nil {
		return nil, fmt.Errorf("context is empty")
	}
	if val, ok := p[key]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("value of %s not in Context", key)
}
