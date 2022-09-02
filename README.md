# Goで環境変数を色んな方法で取得してみよう

## 1. Viperを使った読み込み

```go
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
```

## 2. `go:embed`とgo-tomlでの読み込み

```go
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
```