package main

import (
	"fmt"
	"github.com/elastic/go-ucfg/yaml"
	"os"
)

type AgentConfig struct {
	Root string `config:"root"`
	Port string `port:"port"`
}
type Config struct {
	Agent AgentConfig `config:"agent"`
}

var Agent *AgentConfig

func ParseConfig(filename string) error {
	configContent, err := yaml.NewConfigWithFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found")
		}
		return err
	}
	config := Config{}
	if err := configContent.Unpack(&config); err != nil {
		return err
	}
	Agent = &config.Agent
	return nil
}
