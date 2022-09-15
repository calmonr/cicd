package command

import (
	"github.com/calmonr/cicd/internal/command"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRoot(r command.Runnable, f *pflag.FlagSet) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "playground",
		Short:         "Go CI/CD Playground.",
		Long:          `Go CI/CD Playground.`,
		Version:       "1.0.0",
		RunE:          r,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().AddFlagSet(f)

	return cmd
}
