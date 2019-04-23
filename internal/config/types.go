package config

// Config holds the configuration information for the application
type Config struct {
	StorageDir string `mapstructure:"storage_dir"`

	GitHub struct {
		Endpoint string `mapstructure:"endpoint"`
		Username string `mapstructure:"username"`
		Token    string `mapstructure:"token"`
	} `mapstructure:"github"`
}
