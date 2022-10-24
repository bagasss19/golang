package http

import (
	_ "fico_ar/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func ServerHttp(handler handler) *fiber.App {

	app := fiber.New()

	app.Use(cors.New())

	v1 := app.Group("/api/v1")
	{
		v1.Get("/health", handler.healthHandler.ServiceHealth)
		v1.Get("/ping", handler.healthHandler.Ping)
	}

	ar := v1.Group("/ar")
	{
		ar.Get("/list", handler.arHandler.GetArList)
		ar.Get("/", handler.arHandler.GetOneAr)
		ar.Delete("/", handler.arHandler.DeleteAr)
		ar.Patch("/", handler.arHandler.UpdateAr)
		ar.Put("/", handler.arHandler.UpdateArStatus)
		ar.Get("/company", handler.arHandler.GetAllCompanyCode)
	}

	giro := v1.Group("/giro")
	{
		giro.Get("/list", handler.giroHandler.GetGiroList)
		giro.Get("/", handler.giroHandler.GetOneGiro)
		giro.Post("/", handler.giroHandler.CreateData)
		giro.Delete("/", handler.giroHandler.DeleteGiro)
		giro.Patch("/", handler.giroHandler.UpdateGiro)
	}

	dp := v1.Group("/dp")
	{
		dp.Get("/list", handler.dpHandler.GetDPList)
		dp.Get("/", handler.dpHandler.GetOneDP)
		dp.Post("/", handler.dpHandler.CreateData)
		dp.Delete("/", handler.dpHandler.DeleteDP)
		dp.Patch("/", handler.dpHandler.UpdateDP)
	}

	dp_detail := v1.Group("/dp_detail")
	{
		dp_detail.Post("/", handler.dpHandler.CreateDataDetail)
		dp_detail.Get("/list", handler.dpHandler.GetDPDetailList)
		dp_detail.Get("/:dp_detail_id", handler.dpHandler.GetOneDPDetail)
		dp_detail.Delete("/:dp_detail_id", handler.dpHandler.DeleteDPDetail)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}
