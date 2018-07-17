package main

import (
	"github.com/kadende/kadende-interfaces/pkg/types"
	"github.com/kadende/kadende-interfaces/spi/instance"
)

type plugin string


func (f plugin) Validate(req *types.Any) error {
	return nil
}

func (f plugin) Provision(spec instance.Spec) (*instance.ID, error) {
	fieldId := instance.ID("aldjlajla")
	return &fieldId, nil
}

func (f plugin) Label(id instance.ID, labels map[string]string) (err error) {
	return nil
}

func (f plugin) Destroy(id instance.ID, context instance.Context) (err error) {
	return nil
}

func (f plugin) DescribeInstances(labels map[string]string, properties bool) (instances []instance.Description, err error)  {
	return []instance.Description{}, nil
}

var Plugin plugin
