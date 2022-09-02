package config_test

import (
	"reflect"
	"testing"
	cfg "vipersample/config"
)

func BenchmarkLoadConfigWithEmbed(b *testing.B) {
	b.ResetTimer()
	dbConfig := &cfg.Database{Port: 5432, User: "viper_sample_user", Password: "password", DbName: "viper_development"}
	b.Run("load config from toml", func(b *testing.B) {
		config, _ := cfg.LoadConfigWithEmbed("local")
		if !reflect.DeepEqual(&config.Database, dbConfig) {
			b.Errorf("diff got config: %v, want: %v", config.Database, dbConfig)
		}
	})
}

func BenchmarkLoadConfigJsonWithEmbed(b *testing.B) {
	b.ResetTimer()
	dbConfig := &cfg.Database{Port: 5432, User: "viper_sample_user", Password: "password", DbName: "viper_development"}
	b.Run("load config from json", func(b *testing.B) {
		config, _ := cfg.LoadConfigJsonWithEmbed("local")
		if !reflect.DeepEqual(&config.Database, dbConfig) {
			b.Errorf("diff got config: %v, want: %v", config.Database, dbConfig)
		}
	})
}

func BenchmarkLoadConfigYamlWithEmbed(b *testing.B) {
	b.ResetTimer()
	dbConfig := &cfg.Database{Port: 5432, User: "viper_sample_user", Password: "password", DbName: "viper_development"}
	b.Run("load config from json", func(b *testing.B) {
		config, _ := cfg.LoadConfigYamlWithEmbed("local")
		if !reflect.DeepEqual(&config.Database, dbConfig) {
			b.Errorf("diff got config: %v, want: %v", config.Database, dbConfig)
		}
	})
}
