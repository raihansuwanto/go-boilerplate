package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Args struct {
	ConfigPath string
}

func ProcessArgs(serviceName string, cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet(serviceName, 1)
	f.StringVar(&a.ConfigPath, "c", ".env", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}

// LoadConfig loads the configuration from a .env file
func LoadConfig(args Args, cfg interface{}) error {

	if err := cleanenv.ReadConfig(args.ConfigPath, cfg); err != nil {
		if !os.IsNotExist(err) {
			log.Panic("Failed to read config file")
			return err
		}
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return err
	}

	return nil
}
