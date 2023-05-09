package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Schema   string `yaml:"schema"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		c.Schema,
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
	)
}

func Open(cfg Config) (*pgxpool.Pool, error) {
	// open database connection
	conn, err := pgxpool.New(context.Background(), cfg.DSN())
	if nil != err {
		return nil, err
	}

	// ping database
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	if err = conn.Ping(ctx); nil != err {
		return nil, err
	}

	return conn, nil
}
