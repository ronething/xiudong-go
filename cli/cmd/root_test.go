package cmd

import (
	"runtime"
	"testing"
)

func TestGetRuntime(t *testing.T) {
	t.Log(runtime.GOMAXPROCS(0) + 1) // cpu 核数 + 1
}
