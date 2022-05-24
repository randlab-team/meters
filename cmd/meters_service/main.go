package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/randlab-team/meters/config"
	"github.com/randlab-team/meters/db"
	"github.com/randlab-team/meters/handlers"
	"github.com/randlab-team/meters/repository"
)

func main() {
	initLogger()
	appConfig := initConfig()
	dbConn := initDb(appConfig.DbString)
	meterHandlers := initHandlers(dbConn)

	e := initEcho()
	initMiddlewares(e)

	initRoutes(e, meterHandlers)

	if err := e.Start(":8080"); err != nil {
		log.Fatal().Err(err).Msg("failed to run server")
	}
}

func initLogger() {
	log.Logger = log.With().Caller().Logger()
}

func initConfig() config.AppConfig {
	appConfig, err := config.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init application config")
	}

	return appConfig
}

func initDb(dbString string) *sqlx.DB {
	dbConn, err := db.InitDB(dbString)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init db connection")
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping db")
	}

	return dbConn
}

func initHandlers(dbConn *sqlx.DB) *handlers.Meters {
	metersRepo, err := repository.NewMeters(dbConn)
	if err != nil {
		log.Fatal().Err(err).Msg("filed to init meters repo")
	}

	metersHandler := handlers.NewMeters(
		metersRepo,
	)

	return metersHandler
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	return e
}

func initMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
}

func initRoutes(e *echo.Echo, meterHandlers *handlers.Meters) {
	v1group := e.Group("v1/")
	metersGroup := v1group.Group("meters")

	metersGroup.GET("/", meterHandlers.GetAll)
}
