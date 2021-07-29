package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/naoina/toml"
)

type Config struct {
	Listen struct {
		TCP string
	}
	Log struct {
		Level   string
		Backups int
		Maxsize int64
	}
	Database struct {
		Dsn string
	}
	Security struct {
		CookieSecret string
	}
}

var config *Config

func Init(filename string) error {
	if filename == "" {
		var env = "development"
		for _, name := range []string{"GOLANG_ENV", "ENV"} {
			if s := os.Getenv(name); s != "" {
				env = s
				break
			}
		}
		filename = env + ".toml"
	}

	tomlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	config = new(Config)
	if err = toml.Unmarshal(tomlData, config); err != nil {
		return fmt.Errorf("toml.Decode(%#v) error: %+v", filename, err)
	}

	if filename == "development.toml" {
		fmt.Fprintf(os.Stderr, "%s WAN config.go:131 > running in the development mode.\n", time.Now().Format("15:04:05"))
	}

	return nil
}

func GetConfig() *Config {
	return config
}
