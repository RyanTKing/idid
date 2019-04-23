package git

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RyanTKing/idid/internal/config"
)

var (
	ErrMalformedShortHand = errors.New("malformed shorthand issue/pull link")
	ErrNotFound           = errors.New("no resource in that repository exists with that number")

	client = &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
)

// ExpandLink takes a shorthand github link and turns it into the fully qualified URL
// e.g. org/repo#22 -> https://github.com/org/repo/issues/22
func ExpandLink(shorthand string) (string, error) {
	link, err := makeLink(shorthand)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return "", err
	}
	cfg := config.Get()
	if cfg.GitHub.Username != "" && cfg.GitHub.Token != "" {
		req.SetBasicAuth(cfg.GitHub.Username, cfg.GitHub.Token)
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	return parseGitHubRes(link, res)
}

func makeLink(shorthand string) (string, error) {
	parts := strings.Split(shorthand, "#")
	if len(parts) != 2 {
		return "", ErrMalformedShortHand
	}
	orgRepo := parts[0]
	num, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	cfg := config.Get()
	link := fmt.Sprintf("%s/%s/issues/%d", cfg.GitHub.Endpoint, orgRepo, num)
	return link, nil
}

func parseGitHubRes(link string, res *http.Response) (string, error) {
	if res.StatusCode == http.StatusFound {
		return res.Header.Get("Location"), nil
	}
	if res.StatusCode == http.StatusNotFound {
		return "", ErrNotFound
	}
	if res.StatusCode == http.StatusOK {
		return link, nil
	}

	return "", fmt.Errorf("received code %d from GitHub", res.StatusCode)
}
