package lib

import (
	"crave/shared/configuration"
)

func Start(cntr configuration.IContainer) error {

	if err := startApiServer(cntr); err != nil {
		return err
	}
	if err := startGrpcServer(cntr); err != nil {
		return err
	}
	return nil
}
func startApiServer(cntr configuration.IContainer) error {

	return cntr.DefineRoute()
}

func startGrpcServer(cntr configuration.IContainer) error {

	return cntr.DefineGrpc()
}
