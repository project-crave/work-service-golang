package configuration

import (
	"crave/shared/configuration"
)

func NewContainer() configuration.IContainer {
	return NewHubWorkContainer()
}
