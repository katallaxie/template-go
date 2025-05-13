package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/katallaxie/go-template/internal/cfg"
	"github.com/katallaxie/go-template/pkg/loader"
	"github.com/katallaxie/go-template/pkg/spec"

	"github.com/spf13/cobra"
)

var config = cfg.Default()

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	RootCmd.AddCommand(InitCmd)

	RootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", config.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&config.Force, "force", "f", config.Force, "force overwrite")
	RootCmd.PersistentFlags().StringVarP(&config.File, "config", "c", config.File, "config file")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "g",
	Short: "Go templatating tool",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runRoot(cmd.Context())
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(_ context.Context) error {
	cfg := filepath.Clean(config.File)

	s, err := os.ReadFile(cfg)
	if err != nil {
		return err
	}

	var spec spec.Spec
	if err := spec.UnmarshalYAML(s); err != nil {
		return err
	}

	loader := loader.NewUrlLoader()
	for _, tpl := range spec.Templates {
		if err := loader.Load(tpl, "tmp"); err != nil {
			return err
		}
	}

	return nil
}
