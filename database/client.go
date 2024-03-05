package database

import (
	"context"
	"fmt"
	"github.com/exaring/otelpgx"
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

func OpenWithTracer(config Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.DSN())
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	cfg.MaxConns = 7
	cfg.MinConns = 0
	cfg.MaxConnIdleTime = time.Minute * 30
	cfg.MaxConnLifetime = time.Hour
	cfg.HealthCheckPeriod = time.Minute
	cfg.ConnConfig.Tracer = otelpgx.NewTracer()

	conn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	return conn, nil
}
