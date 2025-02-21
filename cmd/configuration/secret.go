package configuration

import (
	craveConfiguration "crave/shared/configuration"
)

type Secret struct {
	Controller string
}

type Dependency struct {
}

type Variable struct {
	Databese   *craveConfiguration.Database
	Secret     *Secret
	Dependency *Dependency
	Api        *craveConfiguration.Api
}

func NewVariable() *Variable {
	variable := &Variable{}
	variable.Secret = &Secret{
		Controller: "secret variable",
	}
	variable.Api = &craveConfiguration.Api{
		Ip:   "localhost",
		Port: 3000,
	}
	return variable
}
