package http

import (
	"fico_ar/delivery/container"
	"fico_ar/domain/ar"
	downpayment "fico_ar/domain/down_payment"
	"fico_ar/domain/giro"
	"fico_ar/domain/health"
)

type handler struct {
	healthHandler health.HealthHandler
	arHandler     ar.ArHandler
	giroHandler   giro.GiroHandler
	dpHandler     downpayment.DPHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		healthHandler: health.NewHealthHandler(container.HealthFeature),
		arHandler:     ar.NewArHandler(container.ArHandler),
		giroHandler:   giro.NewGiroHandler(container.GiroHandler),
		dpHandler:     downpayment.NewGiroHandler(container.DPHandler),
	}
}
