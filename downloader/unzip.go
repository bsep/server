package downloader

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type extractFunc func(*zip.File, string) // maybe add later

func extractFile(f *zip.File, out_path, rm_path string) {
	file_name := strings.Replace(f.Name, rm_path, "", 1)
	path := filepath.Join(out_path, file_name)
	if DebugMode {
		log.Println("Creating", path)
	}
	err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
	logError(err)

	rc, err := f.Open()
	logError(err)
	if !f.FileInfo().IsDir() {
		fileCopy, err := os.Create(path)
		logError(err)
		_, err = io.Copy(fileCopy, rc)
		logError(err)
		fileCopy.Close()
	}
	rc.Close()
}

func Unzip(zip_path, out_path, rm_path string) error {
	r, err := zip.OpenReader(zip_path)
	if err != nil {
		return err
	}

	for _, f := range r.File {
		extractFile(f, out_path, rm_path)
	}

	if err := r.Close(); err != nil {
		return err
	}

	return nil
}
