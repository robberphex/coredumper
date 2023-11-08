//go:build !linux

package coredumper

import "fmt"

func WriteCoreDump(coredumpPath string) (int, error) {
	return -1, fmt.Errorf("this platform does not support WriteCoreDump")
}
