package database

import (
	"errors"
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Driver       string `yaml:"driver"`
	Debug        bool   `yaml:"debug"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Locale       string `yaml:"locale"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		url.QueryEscape(c.Locale),
	)
}

// Container implement container to make database connection with gorm
type Container interface {
	DatabaseConfig() Config
	SetDatabase(db *gorm.DB)
}

func Provider(arg any) error {
	container, ok := arg.(Container)
	if !ok {
		return errors.New("application doesn't implement database container")
	}

	var cfg = container.DatabaseConfig()
	var dialect = dialectFromCfg(cfg)
	var db, err = openConn(dialect, cfg.Debug)
	if nil != err {
		return err
	}

	container.SetDatabase(db)

	return nil
}

func dialectFromCfg(cfg Config) gorm.Dialector {
	switch cfg.Driver {
	case "mysql":
		return mysql.Open(cfg.DSN())
	case "postgresql":
		return postgres.Open(cfg.DSN())
	default:
		return nil
	}
}

func openConn(dialect gorm.Dialector, debug bool) (*gorm.DB, error) {
	// set log level
	var logMode = logger.Silent
	if debug {
		logMode = logger.Info
	}

	// open db connection
	db, err := gorm.Open(
		dialect, &gorm.Config{
			Logger: logger.Default.LogMode(logMode),
		},
	)
	if nil != err {
		return nil, err
	}

	if err = db.Select(`select 1`).Error; nil != err {
		return nil, err
	}

	return db, nil
}
