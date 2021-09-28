package userAdmin

import (
	"gopkg.in/yaml.v2"
)

type DatadogUser struct {
	Name   string
	Email  string
	ID     string
	Status string
}

func (d *DatadogUser) String() string {
	yml, _ := yaml.Marshal(d)
	return string(yml)
}
