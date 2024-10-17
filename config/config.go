package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"events-api/pkg/config"

	"gopkg.in/yaml.v3"
)

var (
	configuration Config
	once          sync.Once
)

const (
	filePathFormat     = "%s/config/yml/%s.yml"
	productionFileName = "production"
)

type Config struct {
	ProviderXRepository config.RestConfig   `yaml:"provider_x_repository"`
	ProviderXDatabase   config.SQLiteConfig `yaml:"provider_x_sqlite"`
}

func GetConfig() Config {
	once.Do(func() {
		configuration = getConfig()
	})

	return configuration
}

func getConfig() Config {
	filePath := getYamlFilePath()

	configFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("yamlFile.Get err: %v", err)
	}

	conf := Config{}
	if err := yaml.Unmarshal(configFile, &conf); err != nil {
		log.Fatalf("yamlFile.Unmarshal err: %v", err)
	}

	return conf
}

func getYamlFilePath() string {
	basePath, _ := os.Getwd()

	filePath := fmt.Sprintf(filePathFormat, basePath, productionFileName)

	return filePath
}
