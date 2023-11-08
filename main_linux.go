package coredumper

import "github.com/robberphex/coredumper/internal"

func WriteCoreDump(coredumpPath string) (int, error) {
	return internal.WriteCoreDump(coredumpPath)
}
