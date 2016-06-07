package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	r := m.Run()
	os.Exit(r)
}
