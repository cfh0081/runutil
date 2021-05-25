package runutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunFuncName(t *testing.T) {

	expected := `github.com/cfh0081/runutil.TestRunFuncName`
	name := RunFuncName()
	assert.Equal(t, expected, name)
}
