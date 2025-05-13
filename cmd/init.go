package cmd

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/katallaxie/go-template/pkg/spec"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runInit(cmd.Context())
	},
}

func runInit(_ context.Context) error {
	cwd, err := config.Cwd()
	if err != nil {
		return err
	}

	path := filepath.Clean(filepath.Join(cwd, spec.DefaultFileName))

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	s := spec.Example()

	_, err = io.Copy(f, s)
	if err != nil {
		return err
	}

	return nil
}
