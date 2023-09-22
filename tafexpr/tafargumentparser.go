package tafexpr

import (
	"github.com/gclkaze/tafexpr/tafexpr/parser"
	"github.com/gclkaze/tafexpr/tafexpr/variablecontext"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

type TAFArgumentParser struct {
	tokens          []antlr.Token
	symbolicNames   []string
	IntValue        int
	DoubleValue     float64
	OnError         bool
	VariableContext variablecontext.IVariableContext

	lexerErrors  *TAFArgumentErrorListener
	parserErrors *TAFArgumentErrorListener
	ErrorMsgs    []TAFParserArgumentError

	Debug   bool
	IsFloat bool
}

func (ap *TAFArgumentParser) clearErrors() {
	if ap.lexerErrors != nil {
		ap.lexerErrors = nil
	}

	if ap.parserErrors != nil {
		ap.parserErrors = nil
	}
}
func (ap *TAFArgumentParser) CanParse(arg string) bool {

	ap.clearErrors()

	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.OnError = true

	ap.VariableContext = &variablecontext.VariableContext{}

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
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

func (ap *TAFArgumentParser) Parse(arg string) bool {

	ap.clearErrors()

	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.OnError = true

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	if l.OnError {
		ap.OnError = true
		ap.ErrorMsgs = l.ErrorMsgs
		return false
	}

	if l.IsFloat {
		ap.DoubleValue = l.popFloat()
		ap.IsFloat = true

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
	for i := 0; i < len(ap.tokens); i++ {
		t := ap.tokens[i]
		log.Printf("%s (%q)\n", ap.symbolicNames[t.GetTokenType()], t.GetText())
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
