package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Port     int      `yaml:"port" envconfig:"PORT"`
	DBconfig DBconfig `yaml:"db_config" envconfig:"DB_CONFIG"`
	Host     string   `yaml:"host" envconfig:"HOST"`
}

type DBconfig struct {
	DBurl string `yaml:"db_url" envconfig:"DB_URL"`
}

func (c *Config) ReadFromFile(logger *zap.SugaredLogger) error {
	//configPath := "/home/e4t4g/Desktop/URLbackPr/cmd/configs/app.yaml"
	configPath := "./configs/app.yaml"

	data, err := os.ReadFile(configPath)
	if err != nil {
		logger.Errorf("incorrect config file: %s", err)
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		logger.Errorf("unnable to unmarshal file: %s", err)
	}
	return nil
}

func (c *Config) REadFromEnv(logger *zap.SugaredLogger) {
	err := envconfig.Process("", &c)
	if err != nil {
		logger.Fatalf("failed to read config from env: %v", err)

	}
}
