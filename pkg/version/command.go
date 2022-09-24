package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Long:  `Print the version and build information.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(fmt.Sprintf("%s version %s", cmd.Root().Name(), cmd.Root().Version))

			return nil
		},
	}
}
