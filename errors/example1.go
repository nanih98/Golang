package main

import (
	"fmt"
	"ioutil"
)

func main() {
	dir, err := ioutil.TempDir("", "temp")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
}
