package main

import "fmt"

func fail() error {
	return fmt.Errorf("purposely fail: %s", "ooops!")
}

func main() {
	if err := fail(); err != nil {
		panic(err)
	}
}
