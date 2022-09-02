package config

import (
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"strings"
)

var (
	//go:embed environments/local.toml
	Local []byte
	//go:embed environments/local.json
	LocalJson []byte
	//go:embed environments/local.yml
	LocalYaml []byte
	//go:embed environments/development.toml
	Development []byte
)

type Config struct {
	Stage    string `json:"stage" yaml:"stage"`
	Database Database
}

type Database struct {
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user""`
	Password string `json:"password" yaml:"password"`
	DbName   string `json:"dbname" yaml:"dbname"`
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

func LoadConfigJsonWithEmbed(flag string) (*Config, error) {
	var localConfig Config
	err := json.Unmarshal(LocalJson, &localConfig)
	if err != nil {
		return nil, err
	}

	switch flag {
	case "local":
		return &localConfig, nil
	default:
		return nil, errors.New("could not load config")
	}
}

func LoadConfigYamlWithEmbed(flag string) (*Config, error) {
	var localConfig Config
	err := yaml.Unmarshal(LocalYaml, &localConfig)
	if err != nil {
		return nil, err
	}

	switch flag {
	case "local":
		return &localConfig, nil
	default:
		return nil, errors.New("could not load config")
	}
}
