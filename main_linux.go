package coredumper

/*
#cgo CFLAGS: -I${SRCDIR}/include/
#cgo CFLAGS: -I${SRCDIR}/src/
#cgo CXXFLAGS: -I${SRCDIR}/include/
#cgo CXXFLAGS: -I${SRCDIR}/src/

#include <coredumper/coredumper.h>
*/
import "C"

func WriteCoreDump(coredumpPath string) int {
	res := int(C.WriteCoreDump(C.CString(coredumpPath)))
	return res
}
