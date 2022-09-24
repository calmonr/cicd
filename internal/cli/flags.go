package cli

import (
	"github.com/spf13/pflag"
)

const (
	HelpFlag       = "help"
	VersionFlag    = "version"
	ConfigFileFlag = "config-file"
	LogLevelFlag   = "log-level"
)

func Flags() *pflag.FlagSet {
	f := new(pflag.FlagSet)

	f.BoolP(HelpFlag, "h", false, "Show help")
	f.BoolP(VersionFlag, "v", false, "Show version")
	f.StringP(ConfigFileFlag, "c", "", "Path to config file")
	f.StringP(LogLevelFlag, "l", "info", "Minimal allowed log level")

	return f
}
