package main

import (
	"log"
	"strings"

	"github.com/calmonr/cicd/cmd/playground/app"
	"github.com/calmonr/cicd/cmd/playground/app/command"
	"github.com/calmonr/cicd/cmd/playground/app/runnable"
	"github.com/calmonr/cicd/internal/cli"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	name := "playground"

	flags := &pflag.FlagSet{
		SortFlags: true,
		ParseErrorsWhitelist: pflag.ParseErrorsWhitelist{
			UnknownFlags: true,
		},
	}

	flags.Init(name, pflag.ContinueOnError)
	flags.SetInterspersed(true)
	flags.AddFlagSet(cli.Flags())
	flags.AddFlagSet(app.Flags())

	viper.AutomaticEnv()
	viper.SetEnvPrefix(name)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := viper.BindPFlags(flags); err != nil {
		log.Fatalf("could not bind flags: %v", err)
	}

	cmd := command.NewRoot(runnable.NewRoot(), flags)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("could not execute command: %v", err)
	}
}
