package pkg

import (
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	// Conf is config instance
	Conf config
)

type (
	config struct {
		App struct {
			Name string
			Mode string
		}

		Server struct {
			Graceful   bool
			Address    string
			OffTimeout int `toml:"off_timeout"`
		}

		Log struct {
			Level string
			File  string
		}

		IPWhitelist struct {
			Enable   bool
			List     []string
			IPDBFile string
		} `toml:"ipwhitelist"`

		Emitter struct {
			Address   string
			SecretKey string `toml:"secret_key"`
		}

		Channels []channel
	}

	channel struct {
		Prefix string
		Key    string
	}
)

// ServerOffTimeout returns is the duration to wait until killing active requests and stopping the server
func (c config) ServerOffTimeout() time.Duration {
	_default := 5 * time.Second
	if c.Server.OffTimeout < 5 {
		return _default
	}
	return time.Duration(c.Server.OffTimeout) * time.Second
}

// Search and return matched channel
func (c config) GetChannel(path string) (channel, bool) {
	var ch channel
	for _, _ch := range Conf.Channels {
		if strings.HasPrefix(path, _ch.Prefix) {
			return _ch, true
		}
	}
	return ch, false
}

// Initialize Conf
func InitConf(file string) {
	if _, err := toml.DecodeFile(file, &Conf); err != nil {
		panic("loading config file error")
	}
}
