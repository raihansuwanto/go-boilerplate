package postgre

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/raihansuwanto/go-boilerplate/package/config"
)

func MakePostgreDBClient(cfg config.Config) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:         cfg.DB.Host + ":" + cfg.DB.Port,
		User:         cfg.DB.Username,
		Password:     cfg.DB.Password,
		Database:     cfg.DB.Database,
		PoolSize:     cfg.DB.MaxOpenConns,
		PoolTimeout:  time.Duration(cfg.DB.ConnMaxLifetime) * time.Second,
		MinIdleConns: cfg.DB.MinIdleConns,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		MaxRetries:   5,
	})
}
