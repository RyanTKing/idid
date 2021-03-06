package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

const (
	storageDir     = "storage_dir"
	githubEndpoint = "github.endpoint"
	githubUsername = "github.username"
	githubToken    = "github.token"
)

var (
	cfg  *Config
	cfgL sync.Mutex
)

// Get returns the configuration struct populated with values
func Get() *Config {
	cfgL.Lock()
	defer cfgL.Unlock()

	if cfg != nil {
		return cfg
	}

	v := viper.New()
	v.SetEnvPrefix("IDID")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetDefault(storageDir, "$HOME/.local/share/idid")
	v.SetDefault(githubEndpoint, "https://github.com")
	v.SetDefault(githubUsername, "")
	v.SetDefault(githubToken, "")
	v.AutomaticEnv()

	cfg = &Config{}
	err := v.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}

	cfg.StorageDir = os.ExpandEnv(cfg.StorageDir)
	return cfg
}

// Destroy destroys the config so it is reinitialized at next use
func Destroy() {
	cfgL.Lock()
	defer cfgL.Unlock()
	cfg = nil
}
