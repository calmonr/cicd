#!/bin/sh

mkdir -p completions

for sh in bash fish zsh; do
  go run ./cmd/playground/main.go completion $sh >completions/playground.$sh
done
