package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config — конфигурация сервиса автодополнения
type Config struct {
	Port         int    `yaml:"port" env:"AC_PORT"`
	QueriesPath  string `yaml:"queries_path" env:"AC_QUERIES_PATH"`
	TyposPath    string `yaml:"typos_path" env:"AC_TYPOS_PATH"`
	DefaultLimit int    `yaml:"default_limit" env:"AC_DEFAULT_LIMIT"`
	Debug        bool   `yaml:"debug" env:"AC_DEBUG"`
}

// Default возвращает конфиг по умолчанию
func Default() *Config {
	return &Config{
		Port:         8080,
		QueriesPath:  "data/queries.json",
		TyposPath:    "data/typos.json",
		DefaultLimit: 10,
		Debug:        false,
	}
}

// Load загружает конфиг из YAML-файла
func Load(yamlPath string) (*Config, error) {
	cfg := Default()

	if yamlPath != "" {
		data, err := os.ReadFile(yamlPath)
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("read config file: %w", err)
		}
		if err == nil {
			if err := yaml.Unmarshal(data, cfg); err != nil {
				return nil, fmt.Errorf("parse config yaml: %w", err)
			}
		}
	}

	cfg.applyEnv()
	return cfg, nil
}

func (c *Config) applyEnv() {
	if v := os.Getenv("AC_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			c.Port = port
		}
	}
	if v := os.Getenv("AC_QUERIES_PATH"); v != "" {
		c.QueriesPath = v
	}
	if v := os.Getenv("AC_TYPOS_PATH"); v != "" {
		c.TyposPath = v
	}
	if v := os.Getenv("AC_DEFAULT_LIMIT"); v != "" {
		if limit, err := strconv.Atoi(v); err == nil {
			c.DefaultLimit = limit
		}
	}
	if v := os.Getenv("AC_DEBUG"); v != "" {
		c.Debug = strings.EqualFold(v, "true") || v == "1"
	}
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}
