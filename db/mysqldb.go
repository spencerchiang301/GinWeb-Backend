package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

type DBConfig struct {
	UserName     string
	Password     string
	Addr         string
	Port         int
	Database     string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func LoadDBConfig() (*DBConfig, error) {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		return nil, err
	}
	dbConfig := &DBConfig{
		UserName:     cfg.Section("MysqlDB").Key("UserName").String(),
		Password:     cfg.Section("MysqlDB").Key("Password").String(),
		Addr:         cfg.Section("MysqlDB").Key("Addr").String(),
		Port:         cfg.Section("MysqlDB").Key("Port").MustInt(),
		Database:     cfg.Section("MysqlDB").Key("Database").String(),
		MaxLifetime:  cfg.Section("MysqlDB").Key("MaxLifetime").MustInt(),
		MaxOpenConns: cfg.Section("MysqlDB").Key("MaxOpenConns").MustInt(),
		MaxIdleConns: cfg.Section("MysqlDB").Key("MaxIdleConns").MustInt(),
	}
	return dbConfig, nil
}

func InitDB() *sql.DB {
	config, err := LoadDBConfig()
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		config.UserName, config.Password, config.Addr, config.Port, config.Database, "utf8")
	if conn, err := sql.Open("mysql", dsn); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("connect to mysql is succeed")
		conn.SetMaxOpenConns(config.MaxOpenConns)
		conn.SetMaxIdleConns(config.MaxIdleConns)
		conn.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
		return conn
	}
}
