package store

import (
	"os"
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

	err := checkDirectory("test")
	require.NoError(err)
	_, err = os.Stat(config.Get().StorageDir)
	require.NoError(err)
}

func (suite *StoreTestSuite) TestGetDirectory() {
	assert := suite.Assert()
	require := suite.Require()

	dir, err := getDirectory()
	require.NoError(err)
	assert.Equal("test/github_com", dir)
}

func TestStoreTestSuite(t *testing.T) {
	tests := new(StoreTestSuite)
	suite.Run(t, tests)
}
