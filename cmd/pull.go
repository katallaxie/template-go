package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull a template",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runPull(cmd.Context())
	},
}

func runPull(_ context.Context) error {
	return nil
}
