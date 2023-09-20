package tafexpr

import "github.com/antlr4-go/antlr/v4"

type TAFArgumentError struct {
	line, column int
	msg          string
}

func (e *TAFArgumentError) Error() string {
	return e.msg
}

type TAFArgumentErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

func (c *TAFArgumentErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &TAFArgumentError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (c *TAFArgumentErrorListener) RuntimeError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &TAFArgumentError{
		line:   line,
		column: column,
		msg:    msg,
	})
}
