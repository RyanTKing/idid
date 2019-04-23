package store

import (
	"fmt"
	"os"
	"os/user"
	"testing"

	"github.com/RyanTKing/idid/internal/config"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
}

func (suite *StoreTestSuite) SetupSuite() {
	cfg := config.Get()
	cfg.StorageDir = "test"
}

func (suite *StoreTestSuite) TearDownSuite() {
	config.Destroy()
}

func (suite *StoreTestSuite) TearDownTest() {
	require := suite.Require()

	err := os.RemoveAll(config.Get().StorageDir)
	require.NoError(err)
}

func (suite *StoreTestSuite) TestCheckDirectory() {
	require := suite.Require()

	err := checkDirectory("test/github_com")
	require.NoError(err)
	_, err = os.Stat(config.Get().StorageDir)
	require.NoError(err)
}

func (suite *StoreTestSuite) TestGetDirectory() {
	assert := suite.Assert()
	require := suite.Require()

	usr, err := user.Current()
	require.NoError(err)
	expected := fmt.Sprintf("%s/.local/share/idid/github_com", usr.HomeDir)
	dir, err := getDirectory()
	require.NoError(err)
	assert.Equal(expected, dir)
}

func TestStoreTestSuite(t *testing.T) {
	tests := new(StoreTestSuite)
	suite.Run(t, tests)
}
