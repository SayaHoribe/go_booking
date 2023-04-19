package dbrepo

import (
	"database/sql"
	"github.com/SayaHoribe/go_booking/internal/config"
	"github.com/SayaHoribe/go_booking/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatbaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
