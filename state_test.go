package python3

import (
	"runtime"
	"testing"
)

func TestPyThreadState_SetAsyncExc(t *testing.T) {
	runtime.LockOSThread()

	Py_Initialize()
	PyEval_InitThreads()

	gil := PyGILState_Ensure()

	exc := PyErr_NewException("my.error", nil, nil)
	if PyErr_Occurred() != nil {
		t.Fatalf("new exception fail")
	}
	changed := PyThreadState_SetAsyncExc(0, exc)
	if changed != 0 {
		t.Fatalf("why change thread 0")
	}

	PyGILState_Release(gil)
	Py_Finalize()
}
