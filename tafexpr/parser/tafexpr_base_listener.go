// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr

import "github.com/antlr4-go/antlr/v4"

// BaseTafexprListener is a complete listener for a parse tree produced by TafexprParser.
type BaseTafexprListener struct{}

var _ TafexprListener = &BaseTafexprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTafexprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTafexprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTafexprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTafexprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTaf_expression is called when production taf_expression is entered.
func (s *BaseTafexprListener) EnterTaf_expression(ctx *Taf_expressionContext) {}

// ExitTaf_expression is called when production taf_expression is exited.
func (s *BaseTafexprListener) ExitTaf_expression(ctx *Taf_expressionContext) {}

// EnterOrderedEvaluation is called when production OrderedEvaluation is entered.
func (s *BaseTafexprListener) EnterOrderedEvaluation(ctx *OrderedEvaluationContext) {}

// ExitOrderedEvaluation is called when production OrderedEvaluation is exited.
func (s *BaseTafexprListener) ExitOrderedEvaluation(ctx *OrderedEvaluationContext) {}

// EnterHandleVarExpression is called when production HandleVarExpression is entered.
func (s *BaseTafexprListener) EnterHandleVarExpression(ctx *HandleVarExpressionContext) {}

// ExitHandleVarExpression is called when production HandleVarExpression is exited.
func (s *BaseTafexprListener) ExitHandleVarExpression(ctx *HandleVarExpressionContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseTafexprListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseTafexprListener) ExitNumber(ctx *NumberContext) {}

// EnterDoubleValue is called when production DoubleValue is entered.
func (s *BaseTafexprListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production DoubleValue is exited.
func (s *BaseTafexprListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTafexprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTafexprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseTafexprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseTafexprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterIndexExpression is called when production IndexExpression is entered.
func (s *BaseTafexprListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production IndexExpression is exited.
func (s *BaseTafexprListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterVarPath is called when production VarPath is entered.
func (s *BaseTafexprListener) EnterVarPath(ctx *VarPathContext) {}

// ExitVarPath is called when production VarPath is exited.
func (s *BaseTafexprListener) ExitVarPath(ctx *VarPathContext) {}

// EnterVar_expression is called when production var_expression is entered.
func (s *BaseTafexprListener) EnterVar_expression(ctx *Var_expressionContext) {}

// ExitVar_expression is called when production var_expression is exited.
func (s *BaseTafexprListener) ExitVar_expression(ctx *Var_expressionContext) {}

// EnterParenthesisExpression is called when production parenthesisExpression is entered.
func (s *BaseTafexprListener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production parenthesisExpression is exited.
func (s *BaseTafexprListener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}
