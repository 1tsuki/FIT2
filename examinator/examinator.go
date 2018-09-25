package examinator

import (
	"bufio"
	"fmt"
	"github.com/1tsuki/pget"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	urlFormat = "http://web.sfc.keio.ac.jp/~%s/FIT2/%s"
)

type Examinator struct {
	loginIds []string
	pget *pget.Pget
}

func NewExaminator(parallel int, timeout time.Duration, filepath string) (*Examinator, error) {
	students, err := readLines(filepath)
	if err != nil {
		return nil, err
	}

	pget := pget.NewPget(parallel, timeout)

	return &Examinator{
		students,
		pget,
	}, nil
}

func (f *Examinator) Download(filename string, callback func(*url.URL, *http.Response) error) error {
	urls, err := f.getUrls(filename)
	if err != nil {
		return err
	}

	if err := f.pget.WithCallback(urls, callback); err != nil {
		return err
	}

	return nil
}

func (f *Examinator) getUrls(filename string) ([]*url.URL, error) {
	urls := make([]*url.URL, len(f.loginIds))
	for k, v := range f.loginIds {
		urlStr := fmt.Sprintf(urlFormat, v, filename)
		url, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}
		urls[k] = url
	}
	return urls, nil
}

func ExtractLoginId(url *url.URL) (string) {
	path := url.Path
	// TODO improve
	return path[2:10]
}

func ExtractFileName(url *url.URL) string {
	return filepath.Base(url.String())
}

func SaveFile(filepath string, body io.Reader) error {
	// create file
	out, err := os.Create(filepath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed creating file %s", filepath))
	}
	defer out.Close()

	// write response
	_, err = io.Copy(out, body)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed writing file %s", filepath))
	}

	return nil
}

func readLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed loading file %s", filepath))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, errors.Wrap(scanner.Err(), fmt.Sprintf("failed scanning file %s", filepath))
}