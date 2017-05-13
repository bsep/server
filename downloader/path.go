package downloader

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAppPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func GetPluginsPath(joinpath string) string {
	dir, err := GetAppPath()
	if err != nil {
		return "plugins/"
	} else {
		return filepath.Join(dir, "./plugins/", joinpath)
	}
}

func GetFilenameFromURL(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}
