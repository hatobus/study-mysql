package config

import "github.com/kelseyhightower/envconfig"

type MySQLConfig struct {
	Host     string `default:"127.0.0.1"`
	Port     string `default:"3306"`
	DBUser     string `default:"root"`
	Password string `default:"mysql"`
	DataBase string `default:""`
}

func Init() (*MySQLConfig, error) {
	config := &MySQLConfig{}
	err := envconfig.Process("", config)

	if err != nil {
		return &MySQLConfig{}, err
	}

	return config, nil
}
