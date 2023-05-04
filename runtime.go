package runtime

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Env pre-define server environment
type Env int8

const (
	_ Env = iota

	// Sandbox represent sandbox environment
	Sandbox

	// Staging represent staging environment
	Staging

	// Production represent production environment
	Production
)

var envString = map[Env]string{
	Sandbox:    "sandbox",
	Staging:    "staging",
	Production: "production",
}

func (e Env) String() string {
	for k, v := range envString {
		if e == k {
			return v
		}
	}
	return "undefined"
}

// DBConfig handle configuration databases
// database config should have driver & dsn
type DBConfig interface {
	Driver() string
	DSN() string
}

type Runtime struct {
	env Env
	db  *gorm.DB
	log *log.Logger
}

func NewRuntime(env Env) *Runtime {
	return &Runtime{
		env: env,
		log: log.New(os.Stdout, "", 0),
	}
}

// Log returning logger
func (r *Runtime) Log() *log.Logger {
	return r.log
}

// OpenDB initialize database connection
func (r *Runtime) OpenDB(cfg DBConfig) *Runtime {
	var dialect gorm.Dialector
	// open database connection by database driver
	// Gorm supported database driver: mysql, postgresql, sql lite & sql server
	// but we only supported mysql & postgresql for size reason.
	switch cfg.Driver() {
	case "mysql":
		dialect = mysql.Open(cfg.DSN())
	case "postgresql":
		dialect = postgres.Open(cfg.DSN())
	}
	r.db, _ = gorm.Open(dialect)

	return r
}

// DB get database connection
func (r *Runtime) DB() (*gorm.DB, error) {
	// check if database already initialize
	// return error if database not initialize
	if nil != r.db {
		// do ping to database to make sure connection already connected
		db, err := r.db.DB()
		if nil != err {
			return nil, err
		}

		if err = db.Ping(); nil != err {
			return nil, err
		}

		return r.db, nil
	}

	return nil, fmt.Errorf("database connection not initialize")
}
