package config

import (
	"fmt"
	"log"
)

type config struct {
	DBConfig DBConfig
}

// C config instance
var C *config

// DBConfig ...
type DBConfig struct {
	IP string
	Port string
	DBName string
	UserName string
	Password string
	DebugMode bool
}

// ToDsn get db dsn
func (d *DBConfig) ToDsn() (dsn string) {
	if d.DBName == "" { log.Fatal("DBName is None") }
	if d.IP == "" { log.Fatal("IP is None") }
	if d.Port == "" { log.Fatal("Port is None") }
	if d.Password == "" { log.Fatal("Password is None") }
	if d.UserName == "" { log.Fatal("UserName is None") }
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.UserName,
		d.Password,
		d.IP,
		d.Port,
		d.DBName,
	)
	return
}