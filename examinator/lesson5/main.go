package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/1tsuki/FIT2/examinator"
)

var (
	writer  io.Writer
	clock   time.Time
	baseDir string
)

const (
	exitCodeOK = iota
	exitCodeInvalidOption
	exitCodeError
)

func init() {
	writer = os.Stdout
	clock = time.Now()

}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(strArgs []string) int {
	var (
		parallel int
		list     string
	)
	flags := flag.NewFlagSet("lesson2", flag.ContinueOnError)
	flags.IntVar(&parallel, "p", 10, "number of download pipelines")
	flags.StringVar(&list, "s", "../.students", "list of students login ids")
	flags.Parse(strArgs)

	baseDir = examinator.FormatTime(time.Now())
	os.Mkdir(baseDir, os.FileMode(0777))

	ex, err := examinator.NewExaminator(parallel, 30*time.Second, list)
	if err != nil {
		printf("error parsing student list: %v", err)
		return exitCodeInvalidOption
	}

	if err := ex.Download("ex5-2.html", saveFile); err != nil {
		printf("error downloading ex4-1.html: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex5-2.js", saveFile); err != nil {
		printf("error downloading ex4-1.js: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex5-4.html", saveFile); err != nil {
		printf("error downloading ex4-2.html: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex5-4.js", saveFile); err != nil {
		printf("error downloading ex4-2.js: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex5-5.html", saveFile); err != nil {
		printf("error downloading ex4-4.html: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex5-5.js", saveFile); err != nil {
		printf("error downloading ex4-4.js: %v", err)
		return exitCodeError
	}
	return exitCodeOK
}

func saveFile(url *url.URL, resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusNotFound, http.StatusForbidden:
		printf("file did not exist: %s\n", url.String())
		return nil
	default:
		printf("downloaded file %s\n", url.String())
	}

	loginName := examinator.ExtractLoginId(url)
	fileName := examinator.ExtractFileName(url)
	os.Mkdir(filepath.Join(baseDir, loginName), os.FileMode(0777))

	return examinator.SaveFile(filepath.Join(baseDir, loginName, fileName), resp.Body)
}

func printf(format string, a ... interface{}) {
	fmt.Fprintf(writer, format, a...)
}
