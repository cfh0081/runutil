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
