// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TafexprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TafexprLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tafexprlexerLexerInit() {
	staticData := &TafexprLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "", "", "", "'['", "']'",
		"'.'", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "MUL", "DIV", "ADD", "SUB", "INTEGER", "DOUBLE", "WHITESPACE",
		"LBR", "RBR", "CON", "DOLLAR", "NUMBER", "VARIABLE_NAME", "PROP",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "MUL", "DIV", "ADD", "SUB", "INTEGER", "DOUBLE", "WHITESPACE",
		"LBR", "RBR", "CON", "LOWERCASE", "UPPERCASE", "UNDERSCORE", "DOLLAR",
		"SINGLELETTER", "VARIABLECON", "NUMBER", "VARIABLE_NAME", "PROP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 126, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 4, 6, 57, 8, 6, 11, 6, 12, 6, 58, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		5, 7, 66, 8, 7, 10, 7, 12, 7, 69, 9, 7, 1, 7, 1, 7, 4, 7, 73, 8, 7, 11,
		7, 12, 7, 74, 3, 7, 77, 8, 7, 1, 8, 4, 8, 80, 8, 8, 11, 8, 12, 8, 81, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 102, 8, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 113, 8,
		19, 10, 19, 12, 19, 116, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 122,
		8, 20, 10, 20, 12, 20, 125, 9, 20, 0, 0, 21, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 0, 27, 0, 29,
		0, 31, 13, 33, 0, 35, 0, 37, 14, 39, 15, 41, 16, 1, 0, 6, 1, 0, 48, 57,
		1, 0, 48, 48, 1, 0, 49, 57, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 97, 122,
		1, 0, 65, 90, 132, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 43, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 47,
		1, 0, 0, 0, 7, 49, 1, 0, 0, 0, 9, 51, 1, 0, 0, 0, 11, 53, 1, 0, 0, 0, 13,
		56, 1, 0, 0, 0, 15, 76, 1, 0, 0, 0, 17, 79, 1, 0, 0, 0, 19, 85, 1, 0, 0,
		0, 21, 87, 1, 0, 0, 0, 23, 89, 1, 0, 0, 0, 25, 91, 1, 0, 0, 0, 27, 93,
		1, 0, 0, 0, 29, 95, 1, 0, 0, 0, 31, 97, 1, 0, 0, 0, 33, 101, 1, 0, 0, 0,
		35, 103, 1, 0, 0, 0, 37, 105, 1, 0, 0, 0, 39, 107, 1, 0, 0, 0, 41, 117,
		1, 0, 0, 0, 43, 44, 5, 40, 0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 5, 41, 0, 0,
		46, 4, 1, 0, 0, 0, 47, 48, 5, 42, 0, 0, 48, 6, 1, 0, 0, 0, 49, 50, 5, 47,
		0, 0, 50, 8, 1, 0, 0, 0, 51, 52, 5, 43, 0, 0, 52, 10, 1, 0, 0, 0, 53, 54,
		5, 45, 0, 0, 54, 12, 1, 0, 0, 0, 55, 57, 7, 0, 0, 0, 56, 55, 1, 0, 0, 0,
		57, 58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 14, 1,
		0, 0, 0, 60, 61, 7, 1, 0, 0, 61, 62, 5, 46, 0, 0, 62, 77, 7, 1, 0, 0, 63,
		67, 7, 2, 0, 0, 64, 66, 7, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 69, 1, 0, 0,
		0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 70, 72, 5, 46, 0, 0, 71, 73, 7, 0, 0, 0, 72, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1,
		0, 0, 0, 76, 60, 1, 0, 0, 0, 76, 63, 1, 0, 0, 0, 77, 16, 1, 0, 0, 0, 78,
		80, 7, 3, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0,
		0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 6, 8, 0, 0, 84, 18,
		1, 0, 0, 0, 85, 86, 5, 91, 0, 0, 86, 20, 1, 0, 0, 0, 87, 88, 5, 93, 0,
		0, 88, 22, 1, 0, 0, 0, 89, 90, 5, 46, 0, 0, 90, 24, 1, 0, 0, 0, 91, 92,
		7, 4, 0, 0, 92, 26, 1, 0, 0, 0, 93, 94, 7, 5, 0, 0, 94, 28, 1, 0, 0, 0,
		95, 96, 5, 95, 0, 0, 96, 30, 1, 0, 0, 0, 97, 98, 5, 36, 0, 0, 98, 32, 1,
		0, 0, 0, 99, 102, 3, 25, 12, 0, 100, 102, 3, 27, 13, 0, 101, 99, 1, 0,
		0, 0, 101, 100, 1, 0, 0, 0, 102, 34, 1, 0, 0, 0, 103, 104, 5, 46, 0, 0,
		104, 36, 1, 0, 0, 0, 105, 106, 3, 13, 6, 0, 106, 38, 1, 0, 0, 0, 107, 108,
		3, 31, 15, 0, 108, 114, 3, 33, 16, 0, 109, 113, 3, 33, 16, 0, 110, 113,
		3, 13, 6, 0, 111, 113, 3, 29, 14, 0, 112, 109, 1, 0, 0, 0, 112, 110, 1,
		0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 40, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117,
		123, 3, 33, 16, 0, 118, 122, 3, 33, 16, 0, 119, 122, 3, 13, 6, 0, 120,
		122, 3, 29, 14, 0, 121, 118, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 120,
		1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 42, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 11, 0, 58, 67, 74, 76,
		81, 101, 112, 114, 121, 123, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TafexprLexerInit initializes any static state used to implement TafexprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTafexprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TafexprLexerInit() {
	staticData := &TafexprLexerLexerStaticData
	staticData.once.Do(tafexprlexerLexerInit)
}

// NewTafexprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTafexprLexer(input antlr.CharStream) *TafexprLexer {
	TafexprLexerInit()
	l := new(TafexprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TafexprLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Tafexpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TafexprLexer tokens.
const (
	TafexprLexerT__0          = 1
	TafexprLexerT__1          = 2
	TafexprLexerMUL           = 3
	TafexprLexerDIV           = 4
	TafexprLexerADD           = 5
	TafexprLexerSUB           = 6
	TafexprLexerINTEGER       = 7
	TafexprLexerDOUBLE        = 8
	TafexprLexerWHITESPACE    = 9
	TafexprLexerLBR           = 10
	TafexprLexerRBR           = 11
	TafexprLexerCON           = 12
	TafexprLexerDOLLAR        = 13
	TafexprLexerNUMBER        = 14
	TafexprLexerVARIABLE_NAME = 15
	TafexprLexerPROP          = 16
)
