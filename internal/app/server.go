package app

import "fmt"

type Config struct {
	BaseURL      string
	ServerAdress string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	cfg.BaseURL = ":8080"
	cfg.ServerAdress = "http://localhost:8080"
	return cfg, nil
}

// String returns values of Config structure.
func (cfg *Config) String() string {
	return fmt.Sprintf(
		"BASE_URL: %s\n"+
			"SERVER_ADDRESS: %s\n",
		cfg.BaseURL,
		cfg.ServerAdress,
	)
}
