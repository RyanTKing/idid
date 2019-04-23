package store

import (
	"net/http"
	"os"
	"testing"
	"time"

	gock "gopkg.in/h2non/gock.v1"

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
	gock.Off()
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

func (suite *StoreTestSuite) TestReadWrite() {
	assert := suite.Assert()
	require := suite.Require()

	t, err := time.Parse(time.RFC822, "23 Apr 19 17:20 EST")
	require.NoError(err)

	tests := []struct {
		msg    string
		issues []string
	}{
		{"test entry 1", []string{"org1/repo1#1"}},
		{"test entry 2", []string{"org2/repo2#2"}},
		{"test entry 3", []string{"org3/repo3#3", "org3/repo3#4"}},
	}

	cfg := config.Get()
	gock.New(cfg.GitHub.Endpoint).Get("/org1/repo1/issues/1").Reply(http.StatusOK)
	gock.New(cfg.GitHub.Endpoint).Get("/org2/repo2/issues/2").
		Reply(http.StatusFound).AddHeader("Location", "https://github.com/org2/repo2/pull/2")
	gock.New(cfg.GitHub.Endpoint).Get("/org3/repo3/issues/3").Reply(http.StatusOK)
	gock.New(cfg.GitHub.Endpoint).Get("/org3/repo3/issues/4").
		Reply(http.StatusFound).AddHeader("Location", "https://github.com/org3/repo3/pull/4")

	dir, err := getDirectory()
	require.NoError(err)
	for _, tt := range tests {
		err := write(t, dir, tt.msg, tt.issues...)
		require.NoError(err, "For: %s", tt.msg)
	}

	require.True(gock.IsDone())
	t2, err := time.Parse(time.RFC822, "23 Apr 19 17:30 EST")
	require.NoError(err)
	entries := []Entry{
		Entry{
			Msg:     "test entry 1",
			Issues:  []Issue{Issue{Shorthand: "org1/repo1#1", URL: "https://github.com/org1/repo1/issues/1"}},
			Created: t,
		},
		Entry{
			Msg:     "test entry 2",
			Issues:  []Issue{Issue{Shorthand: "org2/repo2#2", URL: "https://github.com/org2/repo2/pull/2"}},
			Created: t,
		},
		Entry{
			Msg: "test entry 3",
			Issues: []Issue{
				Issue{Shorthand: "org3/repo3#3", URL: "https://github.com/org3/repo3/issues/3"},
				Issue{Shorthand: "org3/repo3#4", URL: "https://github.com/org3/repo3/pull/4"},
			},
			Created: t,
		},
	}

	foundEntries, err := read(t2, dir)
	require.NoError(err)
	assert.Equal(entries, foundEntries)
}

func TestStoreTestSuite(t *testing.T) {
	tests := new(StoreTestSuite)
	suite.Run(t, tests)
}
