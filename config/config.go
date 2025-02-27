package config

import (
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		OAuth2   *OAuth2   `mapstructure:"oauth2" validate:"required"`
		State    *State    `mapstructure:"state" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}

	Server struct {
		Port           int           `mapstructure:"port" validate:"required"`
		AllowedOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit      string        `mapstructure:"bodyLimit" validate:"required"`
		TimeOut        time.Duration `mapstructure:"timeout" validate:"required"`
	}

	OAuth2 struct {
		PlayerRedirectUrl string    `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string    `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientId          string    `mapstructure:"clientId" validate:"required"`
		ClientSecret      string    `mapstructure:"clientSecret" validate:"required"`
		Endpoints         endpoints `mapstructure:"endpoints" validate:"required"`
		Scopes            []string  `mapstructure:"scopes" validate:"required"`
		UserInfoUrl       string    `mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl         string    `mapstructure:"revokeUrl" validate:"required"`
	}

	endpoints struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}

	State struct {
		Secret    string        `mapstructure:"secret" validate:"required"`
		ExpiresAt time.Duration `mapstructure:"expiresAt" validate:"required"`
		Issuer    string        `mapstructure:"issuer" validate:"required"`
	}
	Database struct {
		Host     string `mapstructure:"secret" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBname   string `mapstructure:"dbname" validate:"required"`
		SSLmode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validateing := validator.New()

		if err := validateing.Struct(configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}
