package loader

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/katallaxie/go-template/pkg/spec"
)

const (
	defaultBufSize = 1024 * 1024
	defaultTimeout = 10 * time.Second
)

var _ Loader = (*URLLoader)(nil)

// URLLoader is a struct that implements the Loader interface for loading templates from a URL.
type URLLoader struct{}

// NewUrlLoader creates a new UrlLoader instance.
func NewURLLoader() *URLLoader {
	return &URLLoader{}
}

// Load loads data from a URL and returns it as a byte slice.
func (l *URLLoader) Load(ctx context.Context, tpl spec.Template, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tpl.URL, nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: defaultTimeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch URL: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	arc, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}

	for _, f := range arc.File {
		filePath := filepath.Join(path, strings.TrimPrefix(f.Name, tpl.Prefix))

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		for {
			_, err := io.CopyN(dstFile, fileInArchive, defaultBufSize)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				return err
			}
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}
