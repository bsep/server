package editor

import (
	"../downloader"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AddFile(filename string, content string) error {
	custompath := downloader.GetPluginsPath("custom") // Get folder path
	if _, err := os.Stat(custompath); os.IsNotExist(err) {
		os.Mkdir(custompath, 0777)
	}

	addpath := filepath.Join(custompath, filename) // File path for added file (absolute)
	if _, err2 := os.Stat(addpath); err2 == nil {
		return errors.New("File already exists: " + filename)
	}

	return EditFile(addpath, content, true)
}

func badPath(path string) bool {
	return strings.Contains(path, "..")
}

func EditFile(relpath string, content string, absolute bool) error {
	if badPath(relpath) == true {
		return errors.New("Bad path")
	} else {
		if absolute == false {
			relpath = downloader.GetPluginsPath(relpath) // now absolute path
		}
		err := ioutil.WriteFile(relpath, []byte(content), 0644)
		return err
	}
}

func DeleteFile(relpath string) error {
	if badPath(relpath) == true {
		return errors.New("Bad path")
	} else {
		extendedpath := downloader.GetPluginsPath(relpath)
		return os.Remove(extendedpath)
	}
}
