package coredumper

/*
#cgo CFLAGS: -I${SRCDIR}/include/
#cgo CFLAGS: -I${SRCDIR}/src/
#cgo CXXFLAGS: -I${SRCDIR}/include/
#cgo CXXFLAGS: -I${SRCDIR}/src/

#include <coredumper/coredumper.h>
*/
import "C"
import (
	"fmt"
	"os"
)

func main1() {
	fmt.Println("Hello!")
	coredumpPath := fmt.Sprintf("/tmp/coredump-%d", os.Getpid())

	res := WriteCoreDump(coredumpPath)
	fmt.Printf("GetCoreDump-> %d\n", res)

	fmt.Printf("CoreDumped to %s\n", coredumpPath)
}

func WriteCoreDump(coredumpPath string) int {
	res := int(C.WriteCoreDump(C.CString(coredumpPath)))
	return res
}
