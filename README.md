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

2-1. `go:embed`はどの拡張子のファイルが一番早いか？
```go
goos: darwin
goarch: arm64
pkg: vipersample/config
BenchmarkLoadConfigWithEmbed
BenchmarkLoadConfigWithEmbed/load_config_from_toml
BenchmarkLoadConfigWithEmbed/load_config_from_toml-10         	1000000000	         0.0000196 ns/op
BenchmarkLoadConfigJsonWithEmbed
BenchmarkLoadConfigJsonWithEmbed/load_config_from_json
BenchmarkLoadConfigJsonWithEmbed/load_config_from_json-10     	1000000000	         0.0000042 ns/op
BenchmarkLoadConfigYamlWithEmbed
BenchmarkLoadConfigYamlWithEmbed/load_config_from_json
BenchmarkLoadConfigYamlWithEmbed/load_config_from_json-10     	1000000000	         0.0000250 ns/op
```
どれも僅差だが、`json`が一番早かった (0.0000042 ns/op)