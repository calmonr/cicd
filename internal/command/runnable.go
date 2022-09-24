package command

import "github.com/spf13/cobra"

type Runnable func(*cobra.Command, []string) error
