package main

import (
	"./downloader"
	"./editor"
	"./parser"
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
	packurl := c.QueryParam("url")
	parsedurl, urlerr := url.Parse(packurl)
	if urlerr != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(urlerr)})
	}

	packname, tmppath, packerr := downloader.DownloadPack(parsedurl.String())
	if packerr != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(packerr)})
	}

	if debugMode {
		downloader.DebugMode = true
	}

	exterr := downloader.ExtractPack(packname, tmppath)
	if exterr != nil {
		return c.JSON(http.StatusOK, successMsg{false, returnError(exterr)})
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
