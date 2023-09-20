package main

import (
	"log"
	"tafexpr/parser"
	"tafexpr/variablecontext"

	"github.com/antlr4-go/antlr/v4"
)

type TAFArgumentParser struct {
	tokens          []antlr.Token
	symbolicNames   []string
	IntValue        int
	DoubleValue     float64
	OnError         bool
	VariableContext variablecontext.IVariableContext
}

func (ap *TAFArgumentParser) Parse(arg string) bool {
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.OnError = true

	lexerErrors := &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	/*	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		if len(lexerErrors.Errors) != 0 {
			return false
		}
		fmt.Printf("t: %v\n", t)
		ap.tokens = append(ap.tokens, t)
	}*/

	ap.symbolicNames = lexer.SymbolicNames

	if len(lexerErrors.Errors) != 0 {
		log.Printf("Lexer %d errors found\n", len(lexerErrors.Errors))
		for _, e := range lexerErrors.Errors {
			log.Println("\t", e.Error())
		}
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext}

	parserErrors := &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(parserErrors.Errors) != 0 {
		log.Printf("Parser %d errors found\n", len(parserErrors.Errors))
		for _, e := range parserErrors.Errors {
			log.Println("\t", e.Error())
		}
		return false
	}

	if l.IsFloat {
		ap.DoubleValue = l.popFloat()

	} else {
		ap.IntValue = l.pop().(int)
	}
	if l.OnError {
		ap.OnError = true
		return false
	}

	ap.OnError = false
	return true
}

func (ap *TAFArgumentParser) PrintTokens() {
	log.Println(len(ap.tokens))
	for i := 0; i < len(ap.tokens); i++ {
		t := ap.tokens[i]
		log.Printf("%s (%q)\n",
			ap.symbolicNames[t.GetTokenType()], t.GetText())
	}
}

func (ap *TAFArgumentParser) GetTokensLength() int {
	if ap == nil {
		return 0
	}
	return len(ap.tokens)
}

func (ap *TAFArgumentParser) TokenExists(t string, txt string) bool {
	if ap == nil {
		return false
	}

	for i := 0; i < len(ap.tokens); i++ {
		tt := ap.tokens[i]
		if ap.symbolicNames[tt.GetTokenType()] == t && txt == tt.GetText() {
			return true
		}
	}
	return false
}
