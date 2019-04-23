package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	storageDirEnvVar  = "IDID_STORAGE_DIR"
	gitEndpointEnvVar = "IDID_GIT_ENDPOINT"
)

type ConfigTestSuite struct {
	suite.Suite
	cfg *Config
}

func (suite *ConfigTestSuite) TearDownTest() {
	cfgL.Lock()
	defer cfgL.Unlock()
	cfg = nil

	for _, envVar := range []string{storageDirEnvVar, gitEndpointEnvVar} {
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

	assert.Equal("~/.local/share/idid", c.StorageDir)
	assert.Equal("https://github.com", c.GitEndpoint)
}

func (suite *ConfigTestSuite) TestSetFromEnv() {
	assert := suite.Assert()
	require := suite.Require()

	err := os.Setenv(storageDirEnvVar, "/local/share/idid")
	require.NoError(err)
	err = os.Setenv(gitEndpointEnvVar, "https://git.homeserver.com")
	require.NoError(err)

	var c *Config
	require.NotPanics(func() {
		c = Get()
	})

	assert.Equal("/local/share/idid", c.StorageDir)
	assert.Equal("https://git.homeserver.com", c.GitEndpoint)
}

func TestConfigTestSuite(t *testing.T) {
	tests := new(ConfigTestSuite)
	suite.Run(t, tests)
}
