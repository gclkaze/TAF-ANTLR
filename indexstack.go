package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type IndexValue struct {
	Expr  string
	Value int
}

type IndexStack struct {
	value             int
	symbolStartPos    int
	varExpression     string
	evalVarExpression string
	change            bool
	init              bool
	stack             []IndexValue
}

func (ind *IndexStack) SetExpression(v string) {
	if ind.init == true {
		return
	}
	ind.varExpression = v
	ind.evalVarExpression = v
	ind.init = true
}

func (ind IndexStack) IsEmpty() bool {
	return ind.varExpression == "" && ind.evalVarExpression == "" && ind.value == 0 && ind.symbolStartPos == 0
}

func (ind *IndexStack) Push(expr string, val int) {
	ind.stack = append(ind.stack, IndexValue{Expr: expr, Value: val})
	//ind.evalVarExpression = strings.Replace(ind.evalVarExpression, expr, strconv.Itoa(val), 1)
}

func (ind *IndexStack) Clear() {
	ind.varExpression = ""
	ind.evalVarExpression = ""

	ind.value = 0
	ind.symbolStartPos = 0
	ind.change = false
	ind.stack = nil
	ind.init = false
}

func (ind IndexStack) Eval() string {
	log.Println("\n\n===========================================>")
	log.Println("Initially ", ind.varExpression)
	log.Println("Rep ", ind.evalVarExpression)
	log.Println("===========================================>")

	s := ind.EvalExpr(ind.varExpression)
	fmt.Printf("s: %v\n", s)
	return s
}

func (ind IndexStack) EvalExpr(expr string) string {
	if len(ind.stack) == 0 {
		log.Println("stack is empty")
		return expr
	}
	rep := expr
	for i := len(ind.stack) - 1; i >= 0; i-- {
		stored := ind.stack[i]
		rep = strings.Replace(rep, stored.Expr, strconv.Itoa(stored.Value), 1)
	}
	return rep
}
