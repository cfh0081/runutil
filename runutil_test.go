package runutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunFuncName(t *testing.T) {

	name := RunFuncName()
	expected := `github.com/cfh0081/runutil.TestRunFuncName`
	assert.Equal(t, name, expected)
}
