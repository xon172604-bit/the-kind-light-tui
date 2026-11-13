//go:build !windows

package main

// enableANSI is a no-op: macOS and Linux terminals handle ANSI codes natively.
func enableANSI() {}
