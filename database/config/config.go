package config

import "github.com/kelseyhightower/envconfig"

type MySQLConfig struct {
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	User     string `default:"root"`
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
