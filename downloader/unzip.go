package downloader

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

var DebugMode bool = false

func logError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func extractFile(f *zip.File, dest string) {
	path := filepath.Join(dest, f.Name)
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

func unzip(zip_path string, out_path string) error {
	r, err := zip.OpenReader(zip_path)
	if err != nil {
		return err
	}

	for _, f := range r.File {
		extractFile(f, out_path)
	}

	if err := r.Close(); err != nil {
		return err
	}

	return nil
}
