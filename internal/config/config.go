package config

import "fmt"

var (
	conf Config
)

type Config struct {
	TcpAddress       string
	TlsAddress       string
	WebsocketAddress string
	ClusterAddress   string
}

func (cfg *Config) Validate() error {
	if cfg.TcpAddress == "" && cfg.TlsAddress == "" && cfg.WebsocketAddress == "" {
		return fmt.Errorf("invalid listen address")
	}
	return nil
}

func Init(cfg Config) error {
	if err := cfg.Validate(); err != nil {
		return err
	}
	conf = cfg
	return nil
}

func GetConfig() *Config {
	return &conf
}
