package main

import (
	"fmt"
	adHttp "github.com/babon21/ad-storage-service/internal/ad/storage/delivery/http"
	adRepository "github.com/babon21/ad-storage-service/internal/ad/storage/repository/postgres"
	adService "github.com/babon21/ad-storage-service/internal/ad/storage/service"
	"github.com/babon21/ad-storage-service/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := config.Init()

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", conf.Database.Username,
		conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	e := echo.New()
	//middL := middleware.InitMiddleware()
	//e.Use(middL.AccessLogMiddleware)
	roomRepo := adRepository.NewPostgresAdRepository(db)
	roomUsecase := adService.NewAdService(roomRepo)
	adHttp.NewAdHandler(e, roomUsecase)

	log.Fatal().Msg(e.Start(":" + conf.Server.Port).Error())
}
