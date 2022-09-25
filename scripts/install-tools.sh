#!/bin/sh

tools=$(go list -f '{{range .Imports}}{{.}} {{end}}' ./tools/tools.go)

go install $tool
