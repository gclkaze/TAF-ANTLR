package main

import (
	"testing"

	"github.com/gclkaze/tafexpr/variablecontext"

	"github.com/stretchr/testify/assert"
)

func TestPathVariableDeclaration(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("$myVar.  x  .y.z[10].x + 2"))
	assert.Equal(t, 2+2, p.IntValue)
}

func TestPathVariableDeclarationWithExpr(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("2 + $myVar.  x [$y.v.c[$l.m[$m.x[10]]]].y.z + 2"))
	assert.Equal(t, 2+2+2, p.IntValue)
}

func TestPathVariableDeclarationError(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, false, p.Parse("error$myVar.x.y.z"))
}

func TestPathVariableArithmeticsSingle(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.Parse("$myVar.x [$y.v.c[$l.m[$m.x[10]]]].y.z "))
	assert.Equal(t, 2, p.IntValue)
}

func TestPathVariableArithmeticsIndexInteger(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	myVar := "$myVar.x.z[10].y.z"
	assert.Equal(t, true, p.Parse("1 + "+myVar+" + "+myVar+" + 5 + 6"))
	assert.Equal(t, 1+2+2+5+6, p.IntValue)
}

// stress!
func TestPathVariableArithmeticsRecIndexExpr(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	myVar := "$myVar.x .z [$first.x [$second.y.z [$third.w.b]]].y.z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+" + "+myVar+" + 5 + 6"))
	assert.Equal(t, 1+8+2+2+5+6, p.IntValue)
}

func TestPathVariableArithmeticsRecIndexExprWithIndexArithmetics(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	myVar := "$myVar.x .z [$first.x [$second.y.z [$third.w.b + 3] + 2 ] + 1].y.z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+" + "+myVar+" + 5 + 6"))
	assert.Equal(t, 1+8+2+2+5+6, p.IntValue)
}

func TestPathVariableArithmeticsRecIndexExprWithIndexArithmeticsSimple(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	myVar := "$myVar.x .z[ $first.y + 10 ].y.z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+" + 1"))
	assert.Equal(t, 1+8+2+1, p.IntValue)
}

func TestPathVariableArithmeticsRecIndexExprCombo(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	myVar := "$myVar.x[2 + 8].y.z[3 + 9].z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+"  + 5 + 6"))
	assert.Equal(t, 1+8+2+5+6, p.IntValue)
}

func TestPathVariableDeclarationWithMock(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x.y.z[10].x", 60)

	assert.Equal(t, true, p.Parse("$myVar.  x  .y.z[10].x + 2"))
	assert.Equal(t, 60+2, p.IntValue)
}

// stress
func TestPathVariableDeclarationWithExprWithMockSimple(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$m.x[10]", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.a[60].y.z", 160)

	assert.Equal(t, true, p.Parse("$myVar.a [ $m.x[10] ].y.z"))
	assert.Equal(t, 160, p.IntValue)
}

func TestPathVariableDeclarationWithExprWithMockSimpleInt(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.a[60].y.z", 160)

	assert.Equal(t, true, p.Parse("88 + $myVar.a [ 60 ].y.z + 8"))
	assert.Equal(t, 88+160+8, p.IntValue)
}

func TestPathVariableDeclarationWithExprWithMock(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$m.x[10]", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$l.m[65]", 70)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$y.v.c[70]", 80)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x[80].y.z", 160)

	assert.Equal(t, true, p.Parse("2 + $myVar.  x [$y.v.c[$l.m[$m.x[10] + 5]]].y.z + 2"))
	assert.Equal(t, 2+160+2, p.IntValue)
}

func TestPathVariableDeclarationWithExprWithMockFromPhone(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$x", 13)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$d.e[15]", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$c.d[2]", 3)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$a.b[3]", 100)

	assert.Equal(t, true, p.Parse("1 + 5 + $a.b[ $c.d [$d.e [ $x + 2 ]]] + 8"))
	assert.Equal(t, 1+5+100+8, p.IntValue)
}

func TestPathVariableDeclarationWithExprWithMockFromPhoneTwoIndices(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$x", 13)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$d.e[15]", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$c.d[2]", 3)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$a.b[3].x.z[18]", 100)

	assert.Equal(t, true, p.Parse("1 + 5 + $a.b[ $c.d [$d.e [ $x + 2 ]]].x.z[$x + 5] + 8"))
	assert.Equal(t, 1+5+100+8, p.IntValue)
}

// stress
func TestPathVariableArithmeticsRecIndexExprComboWithMock(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.l.x[10].y.z[12].z", 60)

	myVar := "$myVar.l.x[2 + 8].y.z[3 + 9].z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+"  + 5 + 6"))
	assert.Equal(t, 1+8+60+5+6, p.IntValue)
}

func TestPathVariableArithmeticsRecIndexExprComboWithMockWithVars(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.l.x[3].y.z[7].z", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 1)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$k", 3)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$m", 4)

	myVar := "$myVar.l.x[$i + $j].y.z[$k + $m].z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+"  + 5 + 6"))
	assert.Equal(t, 1+8+60+5+6, p.IntValue)
}
func TestPathVariableArithmeticsRecIndexExprComboWithMockAndVarInIndex(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.l.x[11].z", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 3)

	myVar := "$myVar.l.x[$j + 2 + 8 - $i].z"
	assert.Equal(t, true, p.Parse("(1 + 8 + $j + "+myVar+"  + 5 + 6 + 10 ) * ($i*"+myVar+")"))
	assert.Equal(t, (1+8+3+60+5+6+10)*(2*60), p.IntValue)
}

func TestPathVariableArithmeticsRecIndexExprComboWithMockTwiceSimpleNoCalculation(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x[10].y.z[12].z", 60)

	myVar := "$myVar.x[10].y.z[12].z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+"  + 5 + 6 + "+myVar))
	assert.Equal(t, 1+8+60+5+6+60, p.IntValue)
}

func TestPathVariableArithmeticsRecIndexExprComboMock(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x[10].y.z[12].z", 60)
	myVar := "$myVar.x[2 + 8].y.z[3 + 9].z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+"  + 5 + 6"))
	assert.Equal(t, 1+8+60+5+6, p.IntValue)
}

// stress
func TestPathVariableArithmeticsRecIndexExprWithIndexArithmeticsMock(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$third.w.b", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$second.y.z[5]", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$first.x[12]", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x.z[61].y.z", 60)
	myVar := "$myVar.x .z [$first.x [$second.y.z [$third.w.b + 3] + 2 ] + 1].y.z"
	assert.Equal(t, true, p.Parse("1 + 8 + "+myVar+" + "+myVar+" + 5 + 6"))
	assert.Equal(t, 1+8+60+60+5+6, p.IntValue)
}

func TestPathVariableArithmeticsWithDeepIndexAnInteger(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$m.x[10]", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$l.m[2]", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$y.v.c[10]", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x[60].y.z", 60)
	assert.Equal(t, true, p.Parse("$myVar.x [$y.v.c[$l.m[$m.x[10]]]].y.z "))
	assert.Equal(t, 60, p.IntValue)
}

func TestPathVariableArithmeticsWithDeepIndexAnIntegerAndFloatNumber(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$m.x[10]", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$l.m[2]", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$y.v.c[10]", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$myVar.x[60].y.z", 60.05)
	assert.Equal(t, true, p.Parse("$myVar.x [$y.v.c[$l.m[$m.x[10]]]].y.z "))
	assert.Equal(t, 60.05, p.DoubleValue)
}
