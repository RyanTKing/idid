package config

import (
	"fmt"
	"os"
	"os/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	storageDirEnvVar     = "IDID_STORAGE_DIR"
	githubEndpointEnvVar = "IDID_GITHUB_ENDPOINT"
	githubUsernameEnvVar = "IDID_GITHUB_USERNAME"
	githubTokenEnvVar    = "IDID_GITHUB_TOKEN"
)

type ConfigTestSuite struct {
	suite.Suite
	cfg *Config
}

func (suite *ConfigTestSuite) TearDownTest() {
	cfgL.Lock()
	defer cfgL.Unlock()
	cfg = nil

	envVars := []string{storageDirEnvVar, githubEndpointEnvVar, githubUsernameEnvVar, githubTokenEnvVar}
	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}
}

func (suite *ConfigTestSuite) TestDefaults() {
	assert := suite.Assert()
	require := suite.Require()

	var c *Config
	require.NotPanics(func() {
		c = Get()
	})

	usr, err := user.Current()
	require.NoError(err)
	assert.Equal(fmt.Sprintf("%s/.local/share/idid", usr.HomeDir), c.StorageDir)
	assert.Equal("https://github.com", c.GitHub.Endpoint)
	assert.Equal("", c.GitHub.Username)
	assert.Equal("", c.GitHub.Token)
}

func (suite *ConfigTestSuite) TestSetFromEnv() {
	assert := suite.Assert()
	require := suite.Require()

	err := os.Setenv(storageDirEnvVar, "/local/share/idid")
	require.NoError(err)
	err = os.Setenv(githubEndpointEnvVar, "https://git.homeserver.com")
	require.NoError(err)
	err = os.Setenv(githubUsernameEnvVar, "user")
	require.NoError(err)
	err = os.Setenv(githubTokenEnvVar, "token")
	require.NoError(err)

	var c *Config
	require.NotPanics(func() {
		c = Get()
	})

	assert.Equal("/local/share/idid", c.StorageDir)
	assert.Equal("https://git.homeserver.com", c.GitHub.Endpoint)
	assert.Equal("user", c.GitHub.Username)
	assert.Equal("token", c.GitHub.Token)
}

func TestConfigTestSuite(t *testing.T) {
	tests := new(ConfigTestSuite)
	suite.Run(t, tests)
}
