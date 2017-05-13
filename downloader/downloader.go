package downloader

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var DebugMode bool = false

func logError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func Download(url string) (string, error) {
	tmp, err := ioutil.TempFile(os.TempDir(), "pack")
	if err != nil {
		return "", err
	}
	defer tmp.Close()

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	_, err = io.Copy(tmp, res.Body)
	if err != nil {
		return "", err
	}

	return tmp.Name(), nil // Path to temporary file
}

func ExtractPack(packname string, tmppath string) error {
	foldername := strings.TrimSuffix(packname, filepath.Ext(packname)) // Get folder name from plugin pack
	outputpath := GetPluginsPath(foldername)                           // Get output path for pack
	return Unzip(tmppath, outputpath, foldername+"/")
}
