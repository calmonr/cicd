package version

import "fmt"

// nolint: gochecknoglobals
var (
	version = "development"
	commit  = "none"
	date    = "unknown"
)

type Version struct {
	Version, Commit, Date string
}

func (v Version) String() string {
	return fmt.Sprintf("%s, commit %s, built at %s.", v.Version, v.Commit, v.Date)
}

func Get() Version {
	return Version{version, commit, date}
}
