package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = append(os.Args, "--a", "-b")
	main()
}
