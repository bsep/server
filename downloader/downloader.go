package downloader

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadPack(url string) (string, string, error) {
	tokens := strings.Split(url, "/")
	filename := tokens[len(tokens)-1]

	tmp, err := ioutil.TempFile(os.TempDir(), "pack")
	if err != nil {
		return "", "", err
	}
	defer tmp.Close()

	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	_, reterr := io.Copy(tmp, res.Body)
	if reterr != nil {
		return "", "", reterr
	}

	return filename, tmp.Name(), nil // Name of the pack, Path to temporary file
}

const zipcmd string = `7za x "%s" -aoa -r -o"%s" *.xml`

func GetAppPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func GetPluginsPath(joinpath string) string {
	dir, err := GetAppPath()
	if err != nil {
		return ""
	} else {
		return filepath.Join(dir, "./plugins/", joinpath)
	}
}

func ExtractPack(packname string, tmppath string) error {
	foldername := strings.TrimSuffix(packname, filepath.Ext(packname)) // Get folder name from plugin pack
	outputpath := GetPluginsPath(foldername)                           // Get output path for pack
	return unzip(tmppath, outputpath)
}
