package config

import "log"

func MakeConfig() Config {
	cfg := Config{}
	err := cfg.Load("blackswan-server")
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

type Config struct {
	Stage           string                `env:"STAGE"`
	DB              DBConfig              `env-prefix:"DB_"`
	WebService      WebServiceConfig      `env-prefix:"WEBSERVICE_"`
	KeycloakService KeycloakServiceConfig `env-prefix:"KEYCLOAKSERVICE_"`
}

type DBConfig struct {
	Host            string `env:"HOST"`
	Port            string `env:"PORT"`
	Username        string `env:"USERNAME"`
	Password        string `env:"PASSWORD"`
	Database        string `env:"DATABASE"`
	SSLMode         string `env:"SSL_MODE" env-default:"disable"`
	MinIdleConns    int    `env:"MIN_IDLE_CONNS" env-default:"10"`
	MaxOpenConns    int    `env:"MAX_OPEN_CONNS" env-default:"100"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME" env-default:"300"`
	Timeout         int    `env:"TIMEOUT" env-default:"5"`
}

type WebServiceConfig struct {
	Address string `yaml:"address" env:"ADDRESS" env-default:"localhost:8081"`
}

type KeycloakServiceConfig struct {
	Realm        string `env:"REALM"`
	Username     string `env:"USERNAME"`
	Password     string `env:"PASSWORD"`
	BaseUrl      string `env:"BASE_URL"`
	ClientId     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

func (c *Config) Load(serviceName string) error {
	return loadConfig(serviceName, c)
}

func loadConfig(serviceName string, cfg interface{}) error {
	arg := ProcessArgs(serviceName, cfg)
	return LoadConfig(arg, cfg)
}
