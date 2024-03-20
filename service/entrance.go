package service

import "github.com/yanglwd/gateway/service/internal/gateway"

var gatewaserviceinstance gateway.GatewayService

func InitAndServe() error {
	err := gatewaserviceinstance.Init()
	if err != nil {
		return err
	}

	go gatewaserviceinstance.Serve()
	return nil
}

func Exit() {
	gatewaserviceinstance.Exit()
}
