package git

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/ryantking/idid/internal/config"
	"github.com/stretchr/testify/suite"
	gock "gopkg.in/h2non/gock.v1"
)

type GitTestSuite struct {
	suite.Suite
}

func (suite *GitTestSuite) TearDownTest() {
	gock.Off()
}

func (suite *GitTestSuite) TestExpandLink() {
	assert := suite.Assert()
	require := suite.Require()
	tests := []struct {
		shorthand string
		link      string
		err       error
	}{
		{"myorg/myrepo#22", "https://github.com/myorg/myrepo/issues/22", nil},
		{"myorg/myrepo#48", "https://github.com/myorg/myrepo/pull/48", nil},
		{"myorg/myrepo#51", "", ErrNotFound},
		{"myorg/myrepo#1001", "", fmt.Errorf("received code 418 from GitHub")},
		{"myorg/myrepo12", "", ErrMalformedShortHand},
		{"myorg#myrepo#22", "", ErrMalformedShortHand},
		{"myorg/myrepo#foo", "", &strconv.NumError{Func: "Atoi", Num: "foo", Err: strconv.ErrSyntax}},
	}

	cfg := config.Get()
	gock.New(cfg.GitHub.Endpoint).Get("/myorg/myrepo/issues/22").Reply(http.StatusOK)
	gock.New(cfg.GitHub.Endpoint).Get("/myorg/myrepo/issues/48").
		Reply(http.StatusFound).AddHeader("Location", "https://github.com/myorg/myrepo/pull/48")
	gock.New(cfg.GitHub.Endpoint).Get("/myorg/myrepo/issues/51").Reply(http.StatusNotFound)
	gock.New(cfg.GitHub.Endpoint).Get("/myorg/myrepo/issues/1001").Reply(http.StatusTeapot)

	for _, tt := range tests {
		link, err := ExpandLink(tt.shorthand)
		if !assert.Equal(tt.err, err, "For: %s", tt.shorthand) {
			continue
		}
		assert.Equal(tt.link, link, "For: %s", tt.shorthand)
	}

	require.True(gock.IsDone())
}

func TestGitTestSuite(t *testing.T) {
	tests := new(GitTestSuite)
	suite.Run(t, tests)
}
