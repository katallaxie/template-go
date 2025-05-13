package loader

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/katallaxie/go-template/pkg/spec"
)

var _ Loader = (*UrlLoader)(nil)

// UrlLoader is a struct that implements the Loader interface for loading data from a URL.
type UrlLoader struct{}

// NewUrlLoader creates a new UrlLoader instance.
func NewUrlLoader() *UrlLoader {
	return &UrlLoader{}
}

// Load loads data from a URL and returns it as a byte slice.
func (l *UrlLoader) Load(tpl spec.Template, path string) error {
	fmt.Println("loading template from URL: ", tpl.URL)
	resp, err := http.Get(tpl.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch URL: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	arc, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range arc.File {
		filePath := filepath.Join(path, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(path)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", filePath)
		}

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

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}
