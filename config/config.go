package config

import (
	"fmt"
	"os"
	"pet-paradise/log"
	"pet-paradise/model/common"

	"github.com/elastic/go-ucfg/yaml"
)

type DatabaseConfig struct {
	User     string `config:"user" validate:"required"`
	Password string `config:"password" validate:"required"`
	Host     string `config:"host" validate:"required"`
	Port     int    `config:"port"`
	DBName   string `config:"dbname" validate:"required"`
}

type LogConfig struct {
	Dir      string `config:"dir" validate:"required"`
	MaxSize  int64  `config:"max-logger-size"`
	FileName string `config:"filename"`
	LogLevel string `config:"log-level"`
}

type ServerConfig struct {
	Agent []string `config:"agent"`
	API   string   `config:"api"`
}

type Config struct {
	Mysql  DatabaseConfig `config:"mysql"`
	Log    LogConfig      `config:"log"`
	Server ServerConfig   `config:"server"`
}

var Server *ServerConfig

func ParseConfig(filename string) error {
	configContent, err := yaml.NewConfigWithFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found")
		}
		return err
	}
	config := Config{}
	if err := configContent.Unpack(&config); err != nil {
		return err
	}

	db := &config.Mysql
	if err := common.ConfigureMysqlDatabase(db.Host, db.Port, db.User, db.Password, db.DBName); err != nil {
		return err
	}

	logConf := &config.Log
	if err := log.ConfigureLogger(logConf.LogLevel, logConf.Dir, logConf.FileName, logConf.MaxSize); err != nil {
		return err
	} else {
		fmt.Printf("Saving logs at %s\n", logConf.Dir)
	}

	Server = &config.Server

	return nil
}
