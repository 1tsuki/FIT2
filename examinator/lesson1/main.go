package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/1tsuki/pget"
	"github.com/1tsuki/FIT2/examinator"
)

var writer io.Writer

const (
	exitCodeOK = iota
	exitCodeError
)

func init() {
	writer = os.Stdout
}
func main() {
	os.Exit(run(os.Args[1:]))
}
func run(strArgs []string) int {
	var (
		parallel int
		list string
	)
	flags := flag.NewFlagSet("lesson1", flag.ContinueOnError)
	flags.IntVar(&parallel, "p", 10, "number of download pipelines")
	flags.StringVar(&list, "s", "./.students", "list of students login ids")
	flags.Parse(strArgs)

	ex, err := examinator.NewExaminator(list)
	students, err := readLines(list)
	if err != nil {
		printf("error parsing students list")
		return exitCodeError
	}

	ex1HtmlURLs := make([]*url.URL, len(students))
	ex2HtmlURLs := make([]*url.URL, len(students))
	ex1CssURLs := make([]*url.URL, len(students))
	ex2CssURLs := make([]*url.URL, len(students))
	for key, student := range students {
		ex1HtmlURL, err := url.Parse(fmt.Sprintf("http://web.sfc.keio.ac.jp/~%s/ex1-1.html", student))
		if err != nil {
			printf("error parsing html url: %v", err)
			return exitCodeError
		}
		ex1HtmlURLs[key] = ex1HtmlURL

		ex2HtmlURL, err := url.Parse(fmt.Sprintf("http://web.sfc.keio.ac.jp/~%s/ex1-2.html", student))
		if err != nil {
			printf("error parsing html url: %v", err)
			return exitCodeError
		}
		ex2HtmlURLs[key] = ex2HtmlURL


		ex1CssUrl, err := url.Parse(fmt.Sprintf("http://web.sfc.keio.ac.jp/~%s/ex1-1.css", student))
		if err != nil {
			printf("error parsing css url: %v", err)
			return exitCodeError
		}
		ex1CssURLs[key] = ex1CssUrl


		ex2CssUrl, err := url.Parse(fmt.Sprintf("http://web.sfc.keio.ac.jp/~%s/ex1-2.css", student))
		if err != nil {
			printf("error parsing css url: %v", err)
			return exitCodeError
		}
		ex2CssURLs[key] = ex2CssUrl
	}

	downloader := pget.NewPget(10, 60 * time.Second)
	err = downloader.WithCallback(ex1HtmlURLs, checkHTML)
	if err != nil {
		printf("error downloading htmls %v", err)
		return exitCodeError
	}

	err = downloader.WithCallback(ex2HtmlURLs, checkHTML)
	if err != nil {
		printf("error downloading htmls %v", err)
		return exitCodeError
	}

	err = downloader.WithCallback(ex1CssURLs, checkCSS)
	if err != nil {
		printf("error downloading csss %v", err)
		return exitCodeError
	}

	err = downloader.WithCallback(ex2CssURLs, checkCSS)
	if err != nil {
		printf("error downloading csss %v", err)
		return exitCodeError
	}

	return exitCodeOK
}

func checkHTML(url *url.URL, body io.Reader) error {
	printf("downloaded file %s\n", url.String())
	loginName := extractLoginName(url)
	fileName := extractFileName(url)
	os.Mkdir(loginName, os.FileMode(0777))

	save := filepath.Join(loginName, fileName)
	printf("%s\n", save)

	err := saveFile(filepath.Join(loginName, fileName), body)
	if err != nil {
		return err
	}

	validatorURL, err := url.Parse(fmt.Sprintf("https://validator.w3.org/nu/?doc=%s", url.EscapedPath()))
	if err != nil {
		return err
	}

	resp, err := http.Get(validatorURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return saveFile(filepath.Join(loginName, "htmlcheck.html"), resp.Body)
}

func checkCSS(url *url.URL, body io.Reader) error {
	printf("downloaded file %s\n", url.String())
	loginName := extractLoginName(url)
	fileName := extractFileName(url)
	os.Mkdir(loginName, os.FileMode(777))

	err := saveFile(filepath.Join(loginName, fileName), body)
	if err != nil {
		return err
	}

	validatorURLStr := fmt.Sprintf("https://jigsaw.w3.org/css-validator/validator?uri=%s&profile=css3svg&usermedium=all&warning=1&vextwarning=&lang=ja", url.EscapedPath())
	validatorURL, err := url.Parse(validatorURLStr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing url %s", validatorURLStr))
	}

	resp, err := http.Get(validatorURL.String())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed getting url %s", validatorURL.String()))
	}
	defer resp.Body.Close()

	return saveFile(filepath.Join(loginName, "csscheck.html"), resp.Body)
}

func printf(format string, a ... interface{}) {
	fmt.Fprintf(writer, format, a...)
}
