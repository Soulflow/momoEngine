package utils

import "testing"

func TestInfo(t *testing.T) {
	Info("nihao")
	Warn("nihao")
	Fail("nihao")
}
