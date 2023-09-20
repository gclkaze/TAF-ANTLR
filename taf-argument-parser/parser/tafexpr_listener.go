// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr

import "github.com/antlr4-go/antlr/v4"

// TafexprListener is a complete listener for a parse tree produced by TafexprParser.
type TafexprListener interface {
	antlr.ParseTreeListener

	// EnterTaf_expression is called when entering the taf_expression production.
	EnterTaf_expression(c *Taf_expressionContext)

	// EnterOrderedEvaluation is called when entering the OrderedEvaluation production.
	EnterOrderedEvaluation(c *OrderedEvaluationContext)

	// EnterHandleVarExpression is called when entering the HandleVarExpression production.
	EnterHandleVarExpression(c *HandleVarExpressionContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterDoubleValue is called when entering the DoubleValue production.
	EnterDoubleValue(c *DoubleValueContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterIndexExpression is called when entering the IndexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterVarPath is called when entering the VarPath production.
	EnterVarPath(c *VarPathContext)

	// EnterVar_expression is called when entering the var_expression production.
	EnterVar_expression(c *Var_expressionContext)

	// EnterParenthesisExpression is called when entering the parenthesisExpression production.
	EnterParenthesisExpression(c *ParenthesisExpressionContext)

	// ExitTaf_expression is called when exiting the taf_expression production.
	ExitTaf_expression(c *Taf_expressionContext)

	// ExitOrderedEvaluation is called when exiting the OrderedEvaluation production.
	ExitOrderedEvaluation(c *OrderedEvaluationContext)

	// ExitHandleVarExpression is called when exiting the HandleVarExpression production.
	ExitHandleVarExpression(c *HandleVarExpressionContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitDoubleValue is called when exiting the DoubleValue production.
	ExitDoubleValue(c *DoubleValueContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitIndexExpression is called when exiting the IndexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitVarPath is called when exiting the VarPath production.
	ExitVarPath(c *VarPathContext)

	// ExitVar_expression is called when exiting the var_expression production.
	ExitVar_expression(c *Var_expressionContext)

	// ExitParenthesisExpression is called when exiting the parenthesisExpression production.
	ExitParenthesisExpression(c *ParenthesisExpressionContext)
}
