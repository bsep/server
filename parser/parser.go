package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Plugin struct {
	Name, Description, Path string
}

const fieldRegexp string = "<(\\w+:)?(ShortName|Description)>([^<]+)<\\/(\\w+:)?(ShortName|Description)>"

func GetFiles(wpath string) []string {
	files := []string{}
	filepath.Walk(wpath, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			files = append(files, fpath)
		}
		return err
	})
	return files
}

func ParseFile(filepath string) (string, string) {
	contentBytes, _ := ioutil.ReadFile(filepath)
	content := string(contentBytes)
	re := regexp.MustCompile(fieldRegexp)
	matches := re.FindAllStringSubmatch(content, -1)
	var name, desc string
	for _, item := range matches {
		if strings.Contains(item[0], "ShortName") {
			name = item[3]
		}
		if strings.Contains(item[0], "Description") {
			desc = item[3]
		}
	}
	return name, desc
}

func ParseFiles(basepath string, files []string) []Plugin {
	plugins := []Plugin{}
	for _, path := range files {
		name, desc := ParseFile(path)
		relpath, _ := filepath.Rel(basepath, path)
		plugin := Plugin{
			Name:        name,
			Description: desc,
			Path:        relpath,
		}
		plugins = append(plugins, plugin)
	}
	return plugins
}
