package main

import (
	"github.com/bsep/server/downloader"
	"github.com/bsep/server/editor"
	"github.com/bsep/server/parser"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"net/url"
)

func returnError(err error) string {
	return fmt.Sprintf(successFalse, err)
}

func handleAll(c echo.Context) error {
	basepath := downloader.GetPluginsPath("")
	files := parser.GetFiles(basepath)
	parsed := parser.ParseFiles(basepath, files)
	return c.JSON(http.StatusOK, parsed)
}

type successMsg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

const (
	successTrue  = "All is well"
	successFalse = "Error: %v"
)

func handlePack(c echo.Context) error {
	query_url := c.QueryParam("url")
	parsed_url, err := url.Parse(query_url)
	if err != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(err)})
	}

	tmp_path, err := downloader.Download(parsed_url.String())
	if err != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(err)})
	}

	if debugMode {
		downloader.DebugMode = true
	}

	pack_name := downloader.GetFilenameFromURL(parsed_url.String())
	err = downloader.ExtractPack(pack_name, tmp_path)
	if err != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(err)})
	} else {
		return c.JSON(http.StatusOK, successMsg{true, successTrue})
	}
}

func handleAdd(c echo.Context) error {
	filename := c.FormValue("filename")
	content := c.FormValue("content")
	err := editor.AddFile(filename, content)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, successMsg{false, returnError(err)})
	} else {
		return c.JSON(http.StatusOK, successMsg{true, successTrue})
	}
}

func handleEdit(c echo.Context) error {
	relpath := c.FormValue("filepath")
	content := c.FormValue("content")

	fmt.Println("File path:", relpath)
	fmt.Println("Content:", content)

	err := editor.EditFile(relpath, content, false)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, successMsg{false, returnError(err)})
	} else {
		return c.JSON(http.StatusOK, successMsg{true, successTrue})
	}
}

func handleDelete(c echo.Context) error {
	relpath := c.QueryParam("filepath")

	err := editor.DeleteFile(relpath)
	if err != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(err)})
	} else {
		return c.JSON(http.StatusOK, successMsg{true, successTrue})
	}
}
