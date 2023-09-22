package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndeclaredVariable(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupNilVariableContext()
	assert.Equal(t, false, p.Parse("$i + 1 + $d"))
	assert.Equal(t, true, p.OnError)
	assert.Equal(t, 2, len(p.ErrorMsgs))
	assert.Equal(t, "Undeclared variable \"$i\"", p.ErrorMsgs[0].Msg.Error())
	assert.Equal(t, RUNTIME_ERROR, p.ErrorMsgs[0].Type)
	assert.Equal(t, "Undeclared variable \"$d\"", p.ErrorMsgs[1].Msg.Error())
	assert.Equal(t, RUNTIME_ERROR, p.ErrorMsgs[1].Type)
}
