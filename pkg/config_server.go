package pkg

import "fmt"

func NewAppConfigFromViper() (AppConfig, error) {
	ac := AppConfig{
		Db: DbConfig{},
		Server: ServerConfig{
			Port: 8080,
		},
	}

	if err := loadCfg(&ac); err != nil {
		return AppConfig{}, err
	}

	return ac, nil
}

func GenerateAppConfig() (string, error) {
	return generateCfgTemplate(AppConfig{})
}

type AppConfig struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

type DbConfig struct {
	ConnectionString string `mapstructure:"connectionString" validate:"required" env:"GOTRACK_DB_CONNECTION_STRING"`
}

type ServerConfig struct {
	Port int `mapstructure:"port" env:"GOTRACK_SERVER_PORT"`
}

func (sc ServerConfig) GetAddress() string {
	return fmt.Sprintf(":%d", sc.Port)
}
