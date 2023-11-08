package internal

/*
#cgo CFLAGS: -I${SRCDIR}/../include/
#cgo CFLAGS: -I${SRCDIR}/../src/
#cgo CXXFLAGS: -I${SRCDIR}/../include/
#cgo CXXFLAGS: -I${SRCDIR}/../src/

#include <coredumper/coredumper.h>
*/
import "C"
import "fmt"

func WriteCoreDump(coredumpPath string) (int, error) {
	res := int(C.WriteCoreDump(C.CString(coredumpPath)))
	if res != 0 {
		return res, fmt.Errorf("WriteCoreDump returns %d", res)
	} else {
		return res, nil
	}
}
