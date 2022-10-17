package container

import (
	"fico_ar/config"
	ar_feature "fico_ar/domain/ar/feature"
	ar_repository "fico_ar/domain/ar/repository"
	dp_feature "fico_ar/domain/down_payment/feature"
	dp_repository "fico_ar/domain/down_payment/repository"
	giro_feature "fico_ar/domain/giro/feature"
	giro_repository "fico_ar/domain/giro/repository"
	health_feature "fico_ar/domain/health/feature"
	health_repository "fico_ar/domain/health/repository"
	"fico_ar/infrastructure/database"
	"fico_ar/infrastructure/logger"
	"fico_ar/infrastructure/shared/constant"
	"fmt"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	HealthFeature     health_feature.HealthFeature
	ArHandler         ar_feature.ArFeature
	GiroHandler       giro_feature.GiroFeature
	DPHandler         dp_feature.DPFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	logger.InitializeLogger(constant.LOGRUS) // choose which log, ZAP or LOGRUS. Default: LOGRUS

	fmt.Println("Loading database...")
	db, err := database.LoadDatabase(config.Database)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading repository's...")
	healthRepository := health_repository.NewHealthFeature(db)
	arRepository := ar_repository.NewArFeature(db)
	dpRepository := dp_repository.NewDPFeature(db)
	giroRepository := giro_repository.NewGiroFeature(db)

	fmt.Println("Loading feature's...")
	healthFeature := health_feature.NewHealthFeature(config, healthRepository)
	arFeature := ar_feature.NewArFeature(config, arRepository)
	dpFeature := dp_feature.NewDPFeature(config, dpRepository)
	giroFeature := giro_feature.NewGiroFeature(config, giroRepository)

	log.Println("Success!")
	return Container{
		EnvironmentConfig: config,
		HealthFeature:     healthFeature,
		ArHandler:         arFeature,
		DPHandler:         dpFeature,
		GiroHandler:       giroFeature,
	}
}
