package config

import (
	_ "embed"
	"errors"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"strings"
)

var (
	//go:embed environments/local.toml
	Local []byte
	//go:embed environments/development.toml
	Development []byte
)

type Config struct {
	Stage    string
	Database struct {
		Port     int
		User     string
		Password string
		DbName   string
	}
}

func LoadConfigWithViper(flag string) (*Config, error) {
	viper.SetConfigName(flag)
	viper.AddConfigPath("config/environments")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadConfigWithEmbed(flag string) (*Config, error) {
	var localConfig Config
	err := toml.Unmarshal(Local, &localConfig)
	if err != nil {
		return nil, err
	}

	var developmentConfig Config
	err = toml.Unmarshal(Development, &developmentConfig)
	if err != nil {
		return nil, err
	}

	switch flag {
	case "local":
		return &localConfig, nil
	case "development":
		return &developmentConfig, nil
	default:
		return nil, errors.New("could not load config")
	}
}
