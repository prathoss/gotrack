package pkg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotrack/pkg/applogger"
)

func NewDependencyContainer(cfg AppConfig) (DependencyContainer, error) {
	dbConn, err := gorm.Open(postgres.Open(cfg.Db.ConnectionString))
	if err != nil {
		return DependencyContainer{}, err
	}

	return DependencyContainer{
		Db:     dbConn,
		Config: cfg,
	}, nil
}

type DependencyContainer struct {
	Db     *gorm.DB
	Config AppConfig
}

func (d *DependencyContainer) Close() error {
	_ = applogger.Close()
	return nil
}
