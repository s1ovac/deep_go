package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("not found")

type Container struct {
	types map[string]any
}

func NewContainer() *Container {
	return &Container{
		types: make(map[string]any),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	if _, exist := c.types[name]; exist {
		return
	}

	c.types[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, exist := c.types[name]
	if !exist {
		return nil, fmt.Errorf("%w, with name: %s", errNotFound, name)
	}

	return constructor, nil
}
