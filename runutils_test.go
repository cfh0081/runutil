package runutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunFuncName(t *testing.T) {
	expected := `github.com/cfh0081/runutils.TestRunFuncName`
	name := RunFuncName()
	assert.Equal(t, expected, name)
}

func TestRunFileAndLine(t *testing.T) {
	// 需要和实际行号严格对应，否则会报错
	expected := `runutils_test.go#18`
	info := RunFileAndLine()
	assert.Equal(t, expected, info)
}
