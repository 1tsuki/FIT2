package main

import (
	"flag"
	"fmt"
	"github.com/1tsuki/FIT2/examinator"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

var writer io.Writer

const (
	exitCodeOK = iota
	exitCodeInvalidOption
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
	flags.StringVar(&list, "s", "../.students", "list of students login ids")
	flags.Parse(strArgs)

	ex, err := examinator.NewExaminator(parallel, 30 * time.Second, list)
	if err != nil {
		printf("error parsing student list: %v", err)
		return exitCodeInvalidOption
	}

	if err := ex.Download("ex1-1.html", checkHTML); err != nil {
		printf("error downloading ex1-1.html: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex1-2.html", checkHTML); err != nil {
		printf("error downloading ex1-2.html: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex1-1.css", checkCSS); err != nil {
		printf("error downloading ex1-1.css: %v", err)
		return exitCodeError
	}

	if err := ex.Download("ex1-2.css", checkCSS); err != nil {
		printf("error downloading ex1-2.css: %v", err)
		return exitCodeError
	}

	return exitCodeOK
}

func checkHTML(url *url.URL, resp *http.Response) error {
	if resp.StatusCode == http.StatusNotFound {
		printf("file did not exist: %s\n", url.String())
		return nil
	}

	printf("downloaded file %s\n", url.String())

	loginName := examinator.ExtractLoginId(url)
	fileName := examinator.ExtractFileName(url)
	os.Mkdir(loginName, os.FileMode(0777))

	err := examinator.SaveFile(filepath.Join(loginName, fileName), resp.Body)
	if err != nil {
		return err
	}

	validatorURL, err := url.Parse(fmt.Sprintf("https://validator.w3.org/nu/?doc=%s", url.String()))
	if err != nil {
		return err
	}

	resp2, err := http.Get(validatorURL.String())
	if err != nil {
		return err
	}
	defer resp2.Body.Close()

	return examinator.SaveFile(filepath.Join(loginName, fileName + "_htmlcheck.html"), resp2.Body)
}

func checkCSS(url *url.URL, resp *http.Response) error {
	if resp.StatusCode == http.StatusNotFound {
		printf("file did not exist: %s\n", url.String())
		return nil
	}
	printf("downloaded file %s\n", url.String())
	loginName := examinator.ExtractLoginId(url)
	fileName := examinator.ExtractFileName(url)
	os.Mkdir(loginName, os.FileMode(777))

	err := examinator.SaveFile(filepath.Join(loginName, fileName), resp.Body)
	if err != nil {
		return err
	}

	validatorURLStr := fmt.Sprintf("https://jigsaw.w3.org/css-validator/validator?uri=%s&profile=css3svg&usermedium=all&warning=1&vextwarning=&lang=ja", url.String())
	validatorURL, err := url.Parse(validatorURLStr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing url %s", validatorURLStr))
	}

	resp2, err := http.Get(validatorURL.String())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed getting url %s", validatorURL.String()))
	}
	defer resp2.Body.Close()

	return examinator.SaveFile(filepath.Join(loginName, fileName + "_csscheck.html"), resp2.Body)
}

func printf(format string, a ... interface{}) {
	fmt.Fprintf(writer, format, a...)
}
