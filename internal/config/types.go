package config

// Config holds the configuration information for the application
type Config struct {
	StorageDir  string `mapstructure:"storage_dir"`
	GitEndpoint string `mapstructure:"git_endpoint"`
}
