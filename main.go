//go:build !linux

package coredumper

func WriteCoreDump(coredumpPath string) int {
	return -1
}
