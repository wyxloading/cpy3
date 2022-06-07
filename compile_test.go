package python3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPy_CompileString(t *testing.T) {
	Py_Initialize()

	obj := Py_CompileString(`print("abc")`, `main.py`, Py_file_input)
	assert.NotNil(t, obj)
	obj.DecRef()
}
