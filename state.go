package python3

/*
#include "Python.h"
*/
import "C"

//PyThreadState_SetAsyncExc : https://docs.python.org/3/c-api/init.html#c.PyThreadState_SetAsyncExc
func PyThreadState_SetAsyncExc(threadId uint64, exc *PyObject) int {
	return int(C.PyThreadState_SetAsyncExc((C.ulong)(threadId), toc(exc)))
}
