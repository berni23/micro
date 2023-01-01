package main

import (
	"os"
	"testing"
)

//setup testing environment
func test_main(m *testing.M){


	os.Exit(m.Run())
}