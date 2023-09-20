package main

import (
	"log"
	"strconv"
	"strings"
)

type IndexStack struct {
	Expr           string
	Value          int
	SymbolStartPos int
	CurrentPath    string
	Change         bool
}

func (ind IndexStack) IsEmpty() bool {
	return ind.Expr == "" && ind.Value == 0 && ind.SymbolStartPos == 0
}

func (ind *IndexStack) Save(expr string, i int, start int, path string) {
	//fmt.Println("INDEX PUSH")
	ind.Expr = expr
	ind.Value = i
	ind.CurrentPath = path
	ind.SymbolStartPos = start
	//fmt.Printf("ind: %v\n", ind)
}

func (ind *IndexStack) Clear() {
	ind.Expr = ""
	ind.Value = 0
	ind.SymbolStartPos = 0
	ind.Change = false
}

func (ind IndexStack) Eval(expr *string) {
	log.Println("EVAL VAR EXPRESSION")
	//log.Println("Current Path => " + ind.CurrentPath)
	log.Println(ind.Expr, "=>", ind.Value)
	/*	pre := ind.Expr[:ind.SymbolStartPos-1]
		substr := ind.Expr[ind.SymbolStartPos:]
	*/
	substr := strings.Replace(*expr, ind.Expr, strconv.Itoa(ind.Value), 1)
	newPath := strings.Replace(ind.CurrentPath, ind.Expr, strconv.Itoa(ind.Value), 1)

	log.Println("Old Path =>", ind.CurrentPath)
	ind.CurrentPath = newPath
	log.Println("New Path =>", ind.CurrentPath)

	if *expr == substr {
		log.Println("No change!")
		ind.Change = false
	} else {
		ind.Change = true
	}
	*expr = substr
	*expr = newPath
	log.Println("After EVAL =>", *expr)

	//log.Printf("ind: %v\n", ind)
}
