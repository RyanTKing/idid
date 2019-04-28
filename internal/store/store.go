package store

import (
	"fmt"
	"strings"
	"time"

	"os"

	"github.com/ryantking/idid/internal/config"
)

func getPath(now time.Time, dir string) string {
	y, m, d := now.Date()
	return fmt.Sprintf("%s/%2d%02d%02d.json", dir, y, m, d)
}

func getDirectory() (string, error) {
	cfg := config.Get()
	subDir := strings.Split(cfg.GitHub.Endpoint, "//")[1]
	subDir = strings.Replace(subDir, ".", "_", -1)
	dir := fmt.Sprintf("%s/%s", cfg.StorageDir, subDir)
	if err := checkDirectory(dir); err != nil {
		return "", err
	}

	return dir, nil
}

func checkDirectory(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	} else if err != nil {
		return err
	}

	return nil
}
