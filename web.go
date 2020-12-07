package main

import (
	"github.com/bsep/server/downloader"
	"fmt"
	"os"
	"path/filepath"
)

const webURL string = "https://github.com/bsep/server-web/archive/master.zip"

func initWeb() {
	app_path, err := downloader.GetAppPath()
	if err != nil {
		errorAndExit(err, 2)
	}

	if _, err := os.Stat(filepath.Join(app_path, "public/index.html")); os.IsNotExist(err) {

		public_path := filepath.Join(app_path, "public")
		if err := os.Mkdir(public_path, os.ModeDir|os.ModePerm); os.IsPermission(err) {
			errorAndExit(err, 3)
		}

		fmt.Println("Downloading web files")

		tmp_path, err := downloader.Download(webURL)
		if err != nil {
			errorAndExit(err, 4)
		}

		if err := downloader.Unzip(tmp_path, public_path, "server-web-master/"); err != nil {
			errorAndExit(err, 5)
		}

	}
}
