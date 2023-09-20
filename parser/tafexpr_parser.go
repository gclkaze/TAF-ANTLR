// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TafexprParser struct {
	*antlr.BaseParser
}

var TafexprParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tafexprParserInit() {
	staticData := &TafexprParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "", "", "", "'['", "']'",
		"'.'", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "MUL", "DIV", "ADD", "SUB", "INTEGER", "DOUBLE", "WHITESPACE",
		"LBR", "RBR", "CON", "DOLLAR", "NUMBER", "VARIABLE_NAME", "PROP",
	}
	staticData.RuleNames = []string{
		"taf_expression", "expression", "index_expression", "var_path", "var_expression",
		"parenthesisExpression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 70, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 21,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 29, 8, 1, 10, 1, 12, 1,
		32, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 41, 8, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 49, 8, 3, 5, 3, 51, 8, 3, 10, 3,
		12, 3, 54, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 61, 8, 4, 1, 5, 1,
		5, 1, 5, 3, 5, 66, 8, 5, 1, 5, 1, 5, 1, 5, 0, 1, 2, 6, 0, 2, 4, 6, 8, 10,
		0, 2, 1, 0, 3, 4, 1, 0, 5, 6, 74, 0, 12, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0,
		4, 33, 1, 0, 0, 0, 6, 35, 1, 0, 0, 0, 8, 60, 1, 0, 0, 0, 10, 62, 1, 0,
		0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5, 0, 0, 1, 14, 1, 1, 0, 0, 0, 15, 16,
		6, 1, -1, 0, 16, 21, 5, 7, 0, 0, 17, 21, 5, 8, 0, 0, 18, 21, 3, 10, 5,
		0, 19, 21, 3, 8, 4, 0, 20, 15, 1, 0, 0, 0, 20, 17, 1, 0, 0, 0, 20, 18,
		1, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 30, 1, 0, 0, 0, 22, 23, 10, 6, 0, 0,
		23, 24, 7, 0, 0, 0, 24, 29, 3, 2, 1, 7, 25, 26, 10, 5, 0, 0, 26, 27, 7,
		1, 0, 0, 27, 29, 3, 2, 1, 6, 28, 22, 1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 29,
		32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 3, 1, 0, 0,
		0, 32, 30, 1, 0, 0, 0, 33, 34, 3, 2, 1, 0, 34, 5, 1, 0, 0, 0, 35, 40, 5,
		16, 0, 0, 36, 37, 5, 10, 0, 0, 37, 38, 3, 4, 2, 0, 38, 39, 5, 11, 0, 0,
		39, 41, 1, 0, 0, 0, 40, 36, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 52, 1,
		0, 0, 0, 42, 43, 5, 12, 0, 0, 43, 48, 5, 16, 0, 0, 44, 45, 5, 10, 0, 0,
		45, 46, 3, 4, 2, 0, 46, 47, 5, 11, 0, 0, 47, 49, 1, 0, 0, 0, 48, 44, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 42, 1, 0, 0, 0, 51,
		54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 7, 1, 0, 0,
		0, 54, 52, 1, 0, 0, 0, 55, 61, 1, 0, 0, 0, 56, 57, 5, 15, 0, 0, 57, 58,
		5, 12, 0, 0, 58, 61, 3, 6, 3, 0, 59, 61, 5, 15, 0, 0, 60, 55, 1, 0, 0,
		0, 60, 56, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 9, 1, 0, 0, 0, 62, 65, 5,
		1, 0, 0, 63, 66, 3, 10, 5, 0, 64, 66, 3, 2, 1, 0, 65, 63, 1, 0, 0, 0, 65,
		64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 2, 0, 0, 68, 11, 1, 0, 0,
		0, 8, 20, 28, 30, 40, 48, 52, 60, 65,
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

// TafexprParserInit initializes any static state used to implement TafexprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTafexprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TafexprParserInit() {
	staticData := &TafexprParserStaticData
	staticData.once.Do(tafexprParserInit)
}

// NewTafexprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTafexprParser(input antlr.TokenStream) *TafexprParser {
	TafexprParserInit()
	this := new(TafexprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TafexprParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tafexpr.g4"

	return this
}

// TafexprParser tokens.
const (
	TafexprParserEOF           = antlr.TokenEOF
	TafexprParserT__0          = 1
	TafexprParserT__1          = 2
	TafexprParserMUL           = 3
	TafexprParserDIV           = 4
	TafexprParserADD           = 5
	TafexprParserSUB           = 6
	TafexprParserINTEGER       = 7
	TafexprParserDOUBLE        = 8
	TafexprParserWHITESPACE    = 9
	TafexprParserLBR           = 10
	TafexprParserRBR           = 11
	TafexprParserCON           = 12
	TafexprParserDOLLAR        = 13
	TafexprParserNUMBER        = 14
	TafexprParserVARIABLE_NAME = 15
	TafexprParserPROP          = 16
)

// TafexprParser rules.
const (
	TafexprParserRULE_taf_expression        = 0
	TafexprParserRULE_expression            = 1
	TafexprParserRULE_index_expression      = 2
	TafexprParserRULE_var_path              = 3
	TafexprParserRULE_var_expression        = 4
	TafexprParserRULE_parenthesisExpression = 5
)

// ITaf_expressionContext is an interface to support dynamic dispatch.
type ITaf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsTaf_expressionContext differentiates from other interfaces.
	IsTaf_expressionContext()
}

type Taf_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaf_expressionContext() *Taf_expressionContext {
	var p = new(Taf_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_taf_expression
	return p
}

func InitEmptyTaf_expressionContext(p *Taf_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_taf_expression
}

func (*Taf_expressionContext) IsTaf_expressionContext() {}

func NewTaf_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Taf_expressionContext {
	var p = new(Taf_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_taf_expression

	return p
}

func (s *Taf_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Taf_expressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Taf_expressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TafexprParserEOF, 0)
}

func (s *Taf_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Taf_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Taf_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterTaf_expression(s)
	}
}

func (s *Taf_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitTaf_expression(s)
	}
}

func (p *TafexprParser) Taf_expression() (localctx ITaf_expressionContext) {
	localctx = NewTaf_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TafexprParserRULE_taf_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.expression(0)
	}
	{
		p.SetState(13)
		p.Match(TafexprParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrderedEvaluationContext struct {
	ExpressionContext
}

func NewOrderedEvaluationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrderedEvaluationContext {
	var p = new(OrderedEvaluationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrderedEvaluationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderedEvaluationContext) ParenthesisExpression() IParenthesisExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesisExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesisExpressionContext)
}

func (s *OrderedEvaluationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterOrderedEvaluation(s)
	}
}

func (s *OrderedEvaluationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitOrderedEvaluation(s)
	}
}

type HandleVarExpressionContext struct {
	ExpressionContext
}

func NewHandleVarExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleVarExpressionContext {
	var p = new(HandleVarExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleVarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleVarExpressionContext) Var_expression() IVar_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_expressionContext)
}

func (s *HandleVarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleVarExpression(s)
	}
}

func (s *HandleVarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleVarExpression(s)
	}
}

type NumberContext struct {
	ExpressionContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TafexprParserINTEGER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitNumber(s)
	}
}

type DoubleValueContext struct {
	ExpressionContext
}

func NewDoubleValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleValueContext {
	var p = new(DoubleValueContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DoubleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleValueContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(TafexprParserDOUBLE, 0)
}

func (s *DoubleValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterDoubleValue(s)
	}
}

func (s *DoubleValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitDoubleValue(s)
	}
}

type MulDivContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(TafexprParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(TafexprParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(TafexprParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(TafexprParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (p *TafexprParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TafexprParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TafexprParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(16)
			p.Match(TafexprParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDoubleValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(TafexprParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewOrderedEvaluationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.ParenthesisExpression()
		}

	case 4:
		localctx = NewHandleVarExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Var_expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(28)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(23)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TafexprParserMUL || _la == TafexprParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(24)
					p.expression(7)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(26)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TafexprParserADD || _la == TafexprParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(27)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndex_expressionContext is an interface to support dynamic dispatch.
type IIndex_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndex_expressionContext differentiates from other interfaces.
	IsIndex_expressionContext()
}

type Index_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_expressionContext() *Index_expressionContext {
	var p = new(Index_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_index_expression
	return p
}

func InitEmptyIndex_expressionContext(p *Index_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_index_expression
}

func (*Index_expressionContext) IsIndex_expressionContext() {}

func NewIndex_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_expressionContext {
	var p = new(Index_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_index_expression

	return p
}

func (s *Index_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_expressionContext) CopyAll(ctx *Index_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Index_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexExpressionContext struct {
	Index_expressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyIndex_expressionContext(&p.Index_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Index_expressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (p *TafexprParser) Index_expression() (localctx IIndex_expressionContext) {
	localctx = NewIndex_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TafexprParserRULE_index_expression)
	localctx = NewIndexExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_pathContext is an interface to support dynamic dispatch.
type IVar_pathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVar_pathContext differentiates from other interfaces.
	IsVar_pathContext()
}

type Var_pathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_pathContext() *Var_pathContext {
	var p = new(Var_pathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_path
	return p
}

func InitEmptyVar_pathContext(p *Var_pathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_path
}

func (*Var_pathContext) IsVar_pathContext() {}

func NewVar_pathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_pathContext {
	var p = new(Var_pathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_var_path

	return p
}

func (s *Var_pathContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_pathContext) CopyAll(ctx *Var_pathContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Var_pathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_pathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarPathContext struct {
	Var_pathContext
}

func NewVarPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarPathContext {
	var p = new(VarPathContext)

	InitEmptyVar_pathContext(&p.Var_pathContext)
	p.parser = parser
	p.CopyAll(ctx.(*Var_pathContext))

	return p
}

func (s *VarPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarPathContext) AllPROP() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserPROP)
}

func (s *VarPathContext) PROP(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserPROP, i)
}

func (s *VarPathContext) AllLBR() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserLBR)
}

func (s *VarPathContext) LBR(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserLBR, i)
}

func (s *VarPathContext) AllIndex_expression() []IIndex_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndex_expressionContext); ok {
			len++
		}
	}

	tst := make([]IIndex_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndex_expressionContext); ok {
			tst[i] = t.(IIndex_expressionContext)
			i++
		}
	}

	return tst
}

func (s *VarPathContext) Index_expression(i int) IIndex_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_expressionContext)
}

func (s *VarPathContext) AllRBR() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserRBR)
}

func (s *VarPathContext) RBR(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserRBR, i)
}

func (s *VarPathContext) AllCON() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserCON)
}

func (s *VarPathContext) CON(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, i)
}

func (s *VarPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterVarPath(s)
	}
}

func (s *VarPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitVarPath(s)
	}
}

func (p *TafexprParser) Var_path() (localctx IVar_pathContext) {
	localctx = NewVar_pathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TafexprParserRULE_var_path)
	var _alt int

	localctx = NewVarPathContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(TafexprParserPROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(36)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Index_expression()
		}
		{
			p.SetState(38)
			p.Match(TafexprParserRBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(42)
				p.Match(TafexprParserCON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(43)
				p.Match(TafexprParserPROP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(48)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(44)
					p.Match(TafexprParserLBR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(45)
					p.Index_expression()
				}
				{
					p.SetState(46)
					p.Match(TafexprParserRBR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_expressionContext is an interface to support dynamic dispatch.
type IVar_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE_NAME() antlr.TerminalNode
	CON() antlr.TerminalNode
	Var_path() IVar_pathContext

	// IsVar_expressionContext differentiates from other interfaces.
	IsVar_expressionContext()
}

type Var_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_expressionContext() *Var_expressionContext {
	var p = new(Var_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_expression
	return p
}

func InitEmptyVar_expressionContext(p *Var_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_expression
}

func (*Var_expressionContext) IsVar_expressionContext() {}

func NewVar_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_expressionContext {
	var p = new(Var_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_var_expression

	return p
}

func (s *Var_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_expressionContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(TafexprParserVARIABLE_NAME, 0)
}

func (s *Var_expressionContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *Var_expressionContext) Var_path() IVar_pathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_pathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_pathContext)
}

func (s *Var_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterVar_expression(s)
	}
}

func (s *Var_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitVar_expression(s)
	}
}

func (p *TafexprParser) Var_expression() (localctx IVar_expressionContext) {
	localctx = NewVar_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TafexprParserRULE_var_expression)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(TafexprParserVARIABLE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(TafexprParserCON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Var_path()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Match(TafexprParserVARIABLE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParenthesisExpressionContext is an interface to support dynamic dispatch.
type IParenthesisExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParenthesisExpression() IParenthesisExpressionContext
	Expression() IExpressionContext

	// IsParenthesisExpressionContext differentiates from other interfaces.
	IsParenthesisExpressionContext()
}

type ParenthesisExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesisExpressionContext() *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_parenthesisExpression
	return p
}

func InitEmptyParenthesisExpressionContext(p *ParenthesisExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_parenthesisExpression
}

func (*ParenthesisExpressionContext) IsParenthesisExpressionContext() {}

func NewParenthesisExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_parenthesisExpression

	return p
}

func (s *ParenthesisExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenthesisExpressionContext) ParenthesisExpression() IParenthesisExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesisExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesisExpressionContext)
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenthesisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitParenthesisExpression(s)
	}
}

func (p *TafexprParser) ParenthesisExpression() (localctx IParenthesisExpressionContext) {
	localctx = NewParenthesisExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TafexprParserRULE_parenthesisExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(TafexprParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(63)
			p.ParenthesisExpression()
		}

	case 2:
		{
			p.SetState(64)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(67)
		p.Match(TafexprParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TafexprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TafexprParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
