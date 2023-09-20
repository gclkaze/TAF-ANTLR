package main

import (
	"testing"

	"github.com/tafexpr/variablecontext"

	"github.com/stretchr/testify/assert"
)

func SetupVariableContext() variablecontext.IVariableContext {
	v := &variablecontext.VariableContext{}
	v.Init(true)

	var context variablecontext.IVariableContext = v
	return context
}

func SetupMockVariableContext() variablecontext.IVariableContext {
	v := &variablecontext.MockVariableContext{}
	v.Init(true)

	var context variablecontext.IVariableContext = v
	return context
}
func TestVariableDeclaration(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("$myVar"))
	assert.Equal(t, 1, p.IntValue)
}

func TestVariableDeclaration2(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("$myVa2r"))
	assert.Equal(t, 1, p.IntValue)
}

func TestVariableParenthesisDeclaration(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("($myVar)"))
	assert.Equal(t, 1, p.IntValue)
}

func TestVariableArithmetics(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("1 + 2 + ($myVar) + 5 * 3"))
	assert.Equal(t, 1+2+1+5*3, p.IntValue)
}

func TestVariableDoubleArithmetics(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("(100.22 * (200.33 + 3090.11)) - (60.35 + (90/5.2) * ($myVar2 + 2*($myvar + 1))) + $myVar - $myVar"))
	b := almostEqual((100.22*(200.33+3090.11))-(60.35+(90/5.2)*(1+2*(1+1)))+1-1, p.DoubleValue)
	assert.Equal(t, true, b)
}

/*func TestVariableDeclarationWithParenthesis(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("$myVar"))
	assert.Equal(t, 100+200+3090, p.IntValue)
}*/
