package app

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/raihansuwanto/go-boilerplate/app/ent"
	"github.com/raihansuwanto/go-boilerplate/package/config"
	"github.com/raihansuwanto/go-boilerplate/package/logger"
	"github.com/raihansuwanto/go-boilerplate/package/postgre"
	"github.com/raihansuwanto/go-boilerplate/runner"
	"github.com/raihansuwanto/go-boilerplate/runner/web"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg           config.Config
	webRegistries []web.WebModuleRegistry
	ddlModeler    postgre.Modeler
	db            *pg.DB
}

func newApp(
	cfg config.Config,
	webRegistries []web.WebModuleRegistry,
	ddlModeler postgre.Modeler,
	db *pg.DB,
) *App {
	return &App{
		cfg:           cfg,
		webRegistries: webRegistries,
		ddlModeler:    ddlModeler,
		db:            db,
	}
}

func (a App) Run(ctx context.Context) error {

	logger.WithContext(ctx).Info("Starting app")
	logrus.Info("config: ", a.cfg)

	webService := web.NewWebService(web.WithAddress(a.cfg.WebService.Address))

	webService.RegisterModuleRegistry(a.webRegistries...)

	a.MustCreateSchemaIfNotExist()

	return runner.Run(ctx, webService)
}

func (a App) MustCreateSchemaIfNotExist() {
	_, err := a.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		logrus.Error("failed to create extension uuid-ossp: ", err)
	}
	ent.MustCreateSchemaIfNotExist(a.ddlModeler)
}
