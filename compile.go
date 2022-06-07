package python3

/*
#cgo pkg-config: python3
#include "Python.h"
#include "compile.h"
*/
import "C"
import (
	"unsafe"
)

var (
	Py_single_input = int(C.Py_single_input)
	Py_file_input = int(C.Py_file_input)
	Py_eval_input = int(C.Py_eval_input)
)

// Py_CompileString : https://docs.python.org/zh-cn/3.7/c-api/veryhigh.html?highlight=py_compilestringobject#c.Py_CompileString
func Py_CompileString(str, filename string, start int) *PyObject {
	ccommand := C.CString(str)
	defer C.free(unsafe.Pointer(ccommand))

	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	return togo(C.Py_CompileStringExFlags(ccommand, cfilename, C.int(start), nil, C.int(-1)))
}
